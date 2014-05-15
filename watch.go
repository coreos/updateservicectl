package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"text/tabwriter"
	"time"

	"github.com/coreos/updatectl/third_party/code.google.com/p/go-uuid/uuid"
	"github.com/coreos/updatectl/third_party/github.com/codegangsta/cli"
	"github.com/coreos/updatectl/third_party/github.com/coreos/go-omaha/omaha"

	"github.com/coreos/updatectl/client/update/v1"
)

func WatchCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "watch",
			Usage:       "watch [OPTION]... <appID> <groupID> <clientID> <cmd> <args>",
			Description: `Watch for app versions and exec a script`,
			Action:      handle(watch),
			Flags: []cli.Flag{
				cli.IntFlag{"interval, i", 1, "Update polling interval"},
				cli.StringFlag{"version, v", "0.0.0", "Client ID for the watcher"},
			},
		},
	}
}

func fetchVersion(server string, appID string, groupID string, clientID string, version string, debug bool) string {
	client := &http.Client{}

	// TODO: Fill out the OS field correctly based on /etc/os-release
	request := omaha.NewRequest("lsb", "CoreOS", "", "")
	app := request.AddApp(fmt.Sprintf("{%s}", appID), version)
	app.AddUpdateCheck()
	app.MachineID = clientID
	app.BootId = uuid.New()
	app.Track = groupID

	event := app.AddEvent()
	event.Type = "1"
	event.Result = "0"

	raw, err := xml.MarshalIndent(request, "", " ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return version
	}

	if debug {
		fmt.Fprintf(os.Stderr, "Request: %s%s\n", xml.Header, raw)
	}

	resp, err := client.Post(server+"/v1/update/", "text/xml", bytes.NewReader(raw))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return version
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return version
	}

	if debug {
		fmt.Fprintf(os.Stderr, "Response: %s%s\n", xml.Header, string(body))
	}

	oresp := &omaha.Response{}
	err = xml.Unmarshal(body, oresp)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return version
	}

	if oresp.Apps[0].UpdateCheck.Status == "noupdate" {
		return version
	} else if oresp.Apps[0].UpdateCheck.Status == "error-version" {
		return version
	}

	newVersion := oresp.Apps[0].UpdateCheck.Manifest.Version

	return newVersion
}

func runCmd(cmdName string, args []string, version string, oldVersion string, appID string) {
	cmd := exec.Command(cmdName, args...)

	env := os.Environ()
	env = append(env, "UPDATE_SERVICE_VERSION="+version)
	if oldVersion != "" {
		env = append(env, "UPDATE_SERVICE_OLD_VERSION="+oldVersion)
	}
	env = append(env, "UPDATE_SERVICE_APP_ID="+appID)
	cmd.Env = env

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	err = cmd.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)
	cmd.Wait()
}

func watch(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	tick := time.NewTicker(time.Second * time.Duration(c.Int("interval")))
	server := c.GlobalString("server")
	debug := c.GlobalBool("debug")
	version := c.String("version")
	args := c.Args()

	if len(args) < 4 {
		log.Fatalf("appID, groupID and clientID required")
	}

	appID := args[0]
	groupID := args[1]
	clientID := args[2]

	// TODO: Have a better way of asking omaha for whatever version
	version = fetchVersion(server, appID, groupID, clientID, version, debug)
	runCmd(args[3], args[4:], version, "", appID)

	for {
		select {
		case <-tick.C:
			newVersion := fetchVersion(server, appID, groupID, clientID, version, debug)
			if newVersion != version {
				runCmd(args[3], args[4:], newVersion, version, appID)
			}
			version = newVersion
		}
	}

	tick.Stop()
}
