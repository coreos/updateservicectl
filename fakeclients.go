package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"text/tabwriter"
	"time"

	"github.com/coreos-inc/updatectl/third_party/code.google.com/p/go-uuid/uuid"
	"github.com/coreos-inc/updatectl/third_party/github.com/codegangsta/cli"
	"github.com/coreos-inc/updatectl/third_party/github.com/coreos/go-omaha/omaha"

	"github.com/coreos-inc/updatectl/client/update/v1"
)

var verbose bool
var OEM string
var minSleep int

func FakeClientsCommand() []cli.Command {
	return []cli.Command{
		{
			Name:        "fakeclients",
			Usage:       "fakeclients <appid> <groupid> <version>",
			Description: "Simulate multiple clients.",
			Action:      handle(fakeClients),
			Flags: []cli.Flag{
				cli.BoolFlag{"verbose, v", "Print out the request bodies"},
				cli.IntFlag{"clientsperapp, c", 20, "Number of fake clients per appid."},
				cli.IntFlag{"minsleep, m", 1, "Minimum time between update checks."},
				cli.IntFlag{"maxsleep, M", 10, "Maximum time between update checks."},
				cli.IntFlag{"errorrate", 1, "Chance of error (0-100)%."},
				cli.StringFlag{"oem", "fakeclient", "oem to report"},
				// simulate reboot lock.
				cli.IntFlag{"pingonly, p", 0, "halt update and just send ping requests this many times."},
			},
		},
	}
}

type serverConfig struct {
	server string
}

type Client struct {
	Id             string
	SessionId      string
	Version        string
	AppId          string
	Track          string
	config         *serverConfig
	errorRate      int
	pingsRemaining int
}

func (c *Client) Log(format string, v ...interface{}) {
	format = c.Id + ": " + format
	fmt.Printf(format, v...)
}

func (c *Client) OmahaRequest(otype, result string, updateCheck, isPing bool) *omaha.Request {
	req := omaha.NewRequest("lsb", "CoreOS", "", "")
	app := req.AddApp(c.AppId, c.Version)
	app.MachineID = c.Id
	app.BootId = c.SessionId
	app.Track = c.Track
	app.OEM = OEM

	if updateCheck {
		app.AddUpdateCheck()
	}

	if isPing {
		app.AddPing()
		app.Ping.LastReportDays = "1"
		app.Ping.Status = "1"
	}

	if otype != "" {
		event := app.AddEvent()
		event.Type = otype
		event.Result = result
		if result == "0" {
			event.ErrorCode = "2000"
		} else {
			event.ErrorCode = ""
		}
	}

	return req
}

func (c *Client) MakeRequest(otype, result string, updateCheck, isPing bool) (*omaha.Response, error) {
	client := &http.Client{}
	req := c.OmahaRequest(otype, result, updateCheck, isPing)
	raw, err := xml.MarshalIndent(req, "", " ")
	if err != nil {
		return nil, err
	}

	resp, err := client.Post(c.config.server+"/v1/update/", "text/xml", bytes.NewReader(raw))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	oresp := new(omaha.Response)
	err = xml.NewDecoder(resp.Body).Decode(oresp)
	if err != nil {
		return nil, err
	}

	if verbose {
		raw, _ := xml.MarshalIndent(req, "", " ")
		c.Log("request: %s\n", string(raw))
		raw, _ = xml.MarshalIndent(oresp, "", " ")
		c.Log("response: %s\n", string(raw))
	}

	return oresp, nil
}

func (c *Client) SetVersion(resp *omaha.Response) {
	// A field can potentially be nil.
	defer func() {
		if err := recover(); err != nil {
			c.Log("%s: error setting version: %v", c.Id, err)
		}
	}()

	uc := resp.Apps[0].UpdateCheck
	if uc.Status != "ok" {
		c.Log("%s\n", uc.Status)
		return
	}

	randFailRequest := func(eventType, eventResult string) (failed bool, err error) {
		if rand.Intn(100) <= c.errorRate {
			eventType = "3"
			eventResult = "0"
			failed = true
		}
		_, err = c.MakeRequest(eventType, eventResult, false, false)
		return
	}

	requests := [][]string{
		[]string{"13", "1"}, // downloading
		[]string{"14", "1"}, // downloaded
		[]string{"3", "1"},  // installed
	}

	for i, r := range requests {
		if i > 0 {
			time.Sleep(1 * time.Second)
		}
		failed, err := randFailRequest(r[0], r[1])
		if failed {
			log.Printf("failed to update in eventType: %s, eventResult: %s. Retrying.", r[0], r[1])
			time.Sleep(time.Second * time.Duration(minSleep))
			c.MakeRequest(r[0], r[1], false, false)
			return
		}
		if err != nil {
			log.Println(err)
			return
		}
	}

	// simulate reboot lock for a while
	for c.pingsRemaining > 0 {
		c.MakeRequest("", "", false, true)
		c.pingsRemaining--
		time.Sleep(1 * time.Second)
	}

	c.Log("updated from %s to %s\n", c.Version, uc.Manifest.Version)

	c.Version = uc.Manifest.Version

	_, err := c.MakeRequest("3", "2", false, false) // Send complete with new version.
	if err != nil {
		log.Println(err)
	}

	c.SessionId = uuid.New()
}

// Sleep between n and m seconds
func (c *Client) Loop(n, m int) {
	for {
		randSleep(n, m)

		resp, err := c.MakeRequest("3", "2", true, false)
		if err != nil {
			log.Println(err)
			continue
		}
		c.SetVersion(resp)
	}
}

// Sleeps randomly between n and m seconds.
func randSleep(n, m int) {
	r := m
	if m-n > 0 {
		r = rand.Intn(m-n) + n
	}
	time.Sleep(time.Duration(r) * time.Second)
}

func fakeClients(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	args := c.Args()

	if len(args) != 3 {
		log.Fatalf("app, group and initial version required")
	}

	appId := args[0]
	group := args[1]
	version := args[2]

	server := c.GlobalString("server")
	minSleep = c.Int("minsleep")
	maxSleep := c.Int("maxsleep")
	clientsPerApp := c.Int("clientsperapp")
	verbose = c.Bool("verbose")
	errorRate := c.Int("errorrate")
	OEM = c.String("oem")
	pingonly := c.Int("pingonly")

	conf := &serverConfig{
		server: server,
	}

	for i := 0; i < clientsPerApp; i++ {
		c := &Client{
			Id:             fmt.Sprintf("{fake-client-%03d}", i),
			SessionId:      uuid.New(),
			Version:        version,
			AppId:          appId,
			Track:          group,
			config:         conf,
			errorRate:      errorRate,
			pingsRemaining: pingonly,
		}
		go c.Loop(minSleep, maxSleep)
	}

	// run forever
	wait := make(chan bool)
	<-wait
}
