package main

import (
	"fmt"
	"log"
	"text/tabwriter"

	update "github.com/coreos-inc/updatectl/client/update/v1"
)

var (
	clientUpdateFlags struct {
		groupId string
		appId   string
		start   int64
		end     int64
	}

	cmdListClientUpdates = &Command{
		Name:        "list-clientupdates",
		Usage:       "[OPTION]...",
		Description: "Generates a list of client updates.",
		Run:         listClientUpdates,
	}

	cmdListAppVersions = &Command{
		Name:        "list-appversions",
		Usage:       "[OPTION]...",
		Description: "Generates a list of apps/versions with client count.",
		Run:         listAppVersions,
	}
)

func init() {
	cmdListClientUpdates.Flags.StringVar(&clientUpdateFlags.groupId, "group-id", "", "Group id")
	cmdListClientUpdates.Flags.StringVar(&clientUpdateFlags.appId, "app-id", "", "App id")
	cmdListClientUpdates.Flags.Int64Var(&clientUpdateFlags.start, "start", 0, "Start date filter")
	cmdListClientUpdates.Flags.Int64Var(&clientUpdateFlags.end, "end", 0, "End date filter")

	cmdListAppVersions.Flags.Int64Var(&clientUpdateFlags.start, "start", 0, "Start date filter")
	cmdListAppVersions.Flags.Int64Var(&clientUpdateFlags.end, "end", 0, "End date filter")
}

func listClientUpdates(args []string, service *update.Service, out *tabwriter.Writer) int {
	call := service.Clientupdate.List()
	call.DateStart(clientUpdateFlags.start)
	call.DateEnd(clientUpdateFlags.end)
	call.GroupId(clientUpdateFlags.groupId)
	call.AppId(clientUpdateFlags.appId)
	list, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(out, "AppID\tClientID\tVersion\tLastSeen\tGroup\tStatus\tOEM")
	for _, cl := range list.Items {
		fmt.Fprintf(out, "%s\t%s\t%s\t%s\t%s\t%s\t%s\n", cl.AppId,
			cl.ClientId, cl.Version, cl.LastSeen, cl.GroupId,
			cl.Status, cl.Oem)
	}
	out.Flush()
	return OK
}

func listAppVersions(args []string, service *update.Service, out *tabwriter.Writer) int {
	call := service.Appversion.List()

	call.GroupId(clientUpdateFlags.groupId)
	call.AppId(clientUpdateFlags.appId)

	if clientUpdateFlags.start != 0 {
		call.DateStart(clientUpdateFlags.start)
	}

	if clientUpdateFlags.end != 0 {
		call.DateEnd(clientUpdateFlags.end)
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
