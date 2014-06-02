package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"text/tabwriter"
	"time"

	"code.google.com/p/go-uuid/uuid"
	"github.com/coreos/go-omaha/omaha"

	"github.com/coreos-inc/updatectl/client/update/v1"
)

var (
	watchFlags struct {
		interval int
		version  string
	}
	cmdWatch = &Command{
		Name:        "watch",
		Usage:       "[OPTION]... <appID> <groupID> <clientID> <cmd> <args>",
		Description: `Watch for app versions and exec a script`,
		Run:         watch,
	}
)

func init() {
	cmdWatch.Flags.IntVar(&watchFlags.interval, "interval", 1, "Update polling interval")
	cmdWatch.Flags.StringVar(&watchFlags.version, "version", "0.0.0", "Starting version number")
}

func fetchUpdateCheck(server string, appID string, groupID string, clientID string, version string, debug bool) (*omaha.UpdateCheck, error) {
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
		return nil, err
	}

	if debug {
		fmt.Fprintf(os.Stderr, "Request: %s%s\n", xml.Header, raw)
	}

	resp, err := client.Post(server+"/v1/update/", "text/xml", bytes.NewReader(raw))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	if debug {
		fmt.Fprintf(os.Stderr, "Response: %s%s\n", xml.Header, string(body))
	}

	oresp := &omaha.Response{}
	err = xml.Unmarshal(body, oresp)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	return oresp.Apps[0].UpdateCheck, nil
}

func prepareEnvironment(appID string, version string, oldVersion string, updateCheck *omaha.UpdateCheck) []string {
	env := os.Environ()
	env = append(env, "UPDATE_SERVICE_VERSION="+version)
	if oldVersion != "" {
		env = append(env, "UPDATE_SERVICE_OLD_VERSION="+oldVersion)
	}
	env = append(env, "UPDATE_SERVICE_APP_ID="+appID)

	url, err := url.Parse(updateCheck.Urls.Urls[0].CodeBase)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	url.Path = path.Join(url.Path, updateCheck.Manifest.Packages.Packages[0].Name)
	env = append(env, "UPDATE_SERVICE_URL="+url.String())
	return env
}

func runCmd(cmdName string, args []string, appID string, version string, oldVersion string, updateCheck *omaha.UpdateCheck) {
	cmd := exec.Command(cmdName, args...)
	cmd.Env = prepareEnvironment(appID, version, oldVersion, updateCheck)

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

func watch(args []string, service *update.Service, out *tabwriter.Writer) int {
	tick := time.NewTicker(time.Second * time.Duration(watchFlags.interval))
	server := globalFlags.Server
	debug := globalFlags.Debug
	version := watchFlags.version

	if len(args) < 4 {
		return ERROR_USAGE
	}

	appID := args[0]
	groupID := args[1]
	clientID := args[2]

	// initial check
	updateCheck, err := fetchUpdateCheck(server, appID, groupID, clientID, version, debug)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	runCmd(args[3], args[4:], appID, version, "", updateCheck)

	for {
		select {
		case <-tick.C:

			updateCheck, err := fetchUpdateCheck(server, appID, groupID, clientID, version, debug)
			if err != nil {
				log.Printf("warning: update check failed (%v)\n", err)
				continue
			}

			if updateCheck.Status == "noupdate" {
				continue
			} else if updateCheck.Status == "error-version" {
				continue
			}

			newVersion := updateCheck.Manifest.Version

			if newVersion != version {
				runCmd(args[3], args[4:], appID, newVersion, version, updateCheck)
			}
			version = newVersion
		}
	}

	tick.Stop()
	return OK
}
