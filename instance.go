package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/coreos/go-omaha/omaha"
	"github.com/pborman/uuid"

	update "github.com/coreos/updateservicectl/client/update/v1"
)

var (
	instanceFlags struct {
		groupId       StringFlag
		appId         StringFlag
		start         int64
		end           int64
		verbose       bool
		clientsPerApp int
		minSleep      int
		maxSleep      int
		errorRate     int
		OEM           string
		pingOnly      int
		version       string
		forceUpdate   bool
	}

	cmdInstance = &Command{
		Name:    "instance",
		Usage:   "[OPTION]...",
		Summary: "Operations to view instances.",
		Subcommands: []*Command{
			cmdInstanceListUpdates,
			cmdInstanceListAppVersions,
			cmdInstanceFake,
		},
	}

	cmdInstanceListUpdates = &Command{
		Name:        "instance list-updates",
		Usage:       "[OPTION]...",
		Description: "Generates a list of instance updates.",
		Run:         instanceListUpdates,
	}

	cmdInstanceListAppVersions = &Command{
		Name:        "instance list-app-versions",
		Usage:       "[OPTION]...",
		Description: "Generates a list of apps/versions with instance count.",
		Run:         instanceListAppVersions,
	}

	cmdInstanceFake = &Command{
		Name:        "instance fake",
		Usage:       "[OPTION]...",
		Description: "Simulate multiple fake instances.",
		Run:         instanceFake,
	}
)

func init() {
	cmdInstanceListUpdates.Flags.Var(&instanceFlags.groupId, "group-id", "Group id")
	cmdInstanceListUpdates.Flags.Var(&instanceFlags.appId, "app-id", "App id")
	cmdInstanceListUpdates.Flags.Int64Var(&instanceFlags.start, "start", 0, "Start date filter")
	cmdInstanceListUpdates.Flags.Int64Var(&instanceFlags.end, "end", 0, "End date filter")

	cmdInstanceListAppVersions.Flags.Var(&instanceFlags.groupId, "group-id", "Group id")
	cmdInstanceListAppVersions.Flags.Var(&instanceFlags.appId, "app-id", "App id")
	cmdInstanceListAppVersions.Flags.Int64Var(&instanceFlags.start, "start", 0, "Start date filter")
	cmdInstanceListAppVersions.Flags.Int64Var(&instanceFlags.end, "end", 0, "End date filter")

	cmdInstanceFake.Flags.BoolVar(&instanceFlags.verbose, "verbose", false, "Print out the request bodies")
	cmdInstanceFake.Flags.IntVar(&instanceFlags.clientsPerApp, "clients-per-app", 20, "Number of fake fents per appid.")
	cmdInstanceFake.Flags.IntVar(&instanceFlags.minSleep, "min-sleep", 1, "Minimum time between update checks.")
	cmdInstanceFake.Flags.IntVar(&instanceFlags.maxSleep, "max-sleep", 10, "Maximum time between update checks.")
	cmdInstanceFake.Flags.IntVar(&instanceFlags.errorRate, "errorrate", 1, "Chance of error (0-100)%.")
	cmdInstanceFake.Flags.StringVar(&instanceFlags.OEM, "oem", "fakeclient", "oem to report")
	// simulate reboot lock.
	cmdInstanceFake.Flags.IntVar(&instanceFlags.pingOnly, "ping-only", 0, "halt update and just send ping requests this many times.")
	cmdInstanceFake.Flags.Var(&instanceFlags.appId, "app-id", "Application ID to update.")
	instanceFlags.appId.required = true
	cmdInstanceFake.Flags.Var(&instanceFlags.groupId, "group-id", "Group ID to update.")
	instanceFlags.groupId.required = true
	cmdInstanceFake.Flags.StringVar(&instanceFlags.version, "version", "0.0.0", "Version to report.")
	cmdInstanceFake.Flags.BoolVar(&instanceFlags.forceUpdate, "force-update", false, "Force updates regardless of rate limiting")
}

func instanceListUpdates(args []string, service *update.Service, out *tabwriter.Writer) int {
	call := service.Clientupdate.List()
	call.DateStart(instanceFlags.start)
	call.DateEnd(instanceFlags.end)
	if instanceFlags.groupId.Get() != nil {
		call.GroupId(instanceFlags.groupId.String())
	}
	if instanceFlags.groupId.Get() != nil {
		call.AppId(instanceFlags.appId.String())
	}
	list, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(out, "AppID\tClientID\tVersion\tLastSeen\tGroup\tOEM")
	for _, cl := range list.Items {
		fmt.Fprintf(out, "%s\t%s\t%s\t%s\t%s\t%s\n", cl.AppId,
			cl.ClientId, cl.Version, cl.LastSeen, cl.GroupId, cl.Oem)
	}
	out.Flush()
	return OK
}

func instanceListAppVersions(args []string, service *update.Service, out *tabwriter.Writer) int {
	call := service.Appversion.List()

	if instanceFlags.groupId.Get() != nil {
		call.GroupId(instanceFlags.groupId.String())
	}
	if instanceFlags.appId.Get() != nil {
		call.AppId(instanceFlags.appId.String())
	}
	if instanceFlags.start != 0 {
		call.DateStart(instanceFlags.start)
	}

	if instanceFlags.end != 0 {
		call.DateEnd(instanceFlags.end)
	}

	list, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(out, "AppID\tGroupID\tVersion\tClients")
	for _, cl := range list.Items {
		fmt.Fprintf(out, "%s\t%s\t%s\t%d\n", cl.AppId, cl.GroupId, cl.Version, cl.Count)
	}
	out.Flush()
	return OK
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
	forceUpdate    bool
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
	app.OEM = instanceFlags.OEM
	if c.forceUpdate {
		req.InstallSource = "ondemandupdate"
	} else {
		req.InstallSource = "scheduler"
	}

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

	if instanceFlags.verbose {
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
			c.Log("%s: error setting version: %v\n", c.Id, err)
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
			time.Sleep(time.Second * time.Duration(instanceFlags.minSleep))
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

func randomHex(n int) string {
	rand.Seed(time.Now().UnixNano())

	chars := "abcdef0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func instanceFake(args []string, service *update.Service, out *tabwriter.Writer) int {
	if instanceFlags.appId.Get() == nil || instanceFlags.groupId.Get() == nil {
		return ERROR_USAGE
	}

	conf := &serverConfig{
		server: globalFlags.Server,
	}

	// generate a prefix with a well-known string and a constant sequence of hex
	// this lets us easily recognize fake instances
	// it still has to be a valid uuid though
	prefix := "deadbeef" + randomHex(6)

	for i := 0; i < instanceFlags.clientsPerApp; i++ {
		c := &Client{
			Id:             prefix + strings.Replace(uuid.New(), "-", "", -1)[14:],
			SessionId:      uuid.New(),
			Version:        instanceFlags.version,
			AppId:          instanceFlags.appId.String(),
			Track:          instanceFlags.groupId.String(),
			config:         conf,
			errorRate:      instanceFlags.errorRate,
			pingsRemaining: instanceFlags.pingOnly,
			forceUpdate:    instanceFlags.forceUpdate,
		}
		go c.Loop(instanceFlags.minSleep, instanceFlags.maxSleep)
	}

	// run forever
	wait := make(chan bool)
	<-wait
	return OK
}
