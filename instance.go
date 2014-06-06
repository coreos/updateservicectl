package main

import (
	"fmt"
	"log"
	"text/tabwriter"

	update "github.com/coreos-inc/updatectl/client/update/v1"
)

var (
	instanceFlags struct {
		groupId StringFlag
		appId   StringFlag
		start   int64
		end     int64
	}
	cmdInstance = &Command{
		Name:        "instance",
		Usage:       "[OPTION]...",
		Description: "Operations to view instances.",
		Subcommands: []*Command{
			cmdInstanceListUpdates,
			cmdInstanceListAppVersions,
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
}

func instanceListUpdates(args []string, service *update.Service, out *tabwriter.Writer) int {
	if instanceFlags.appId.Get() == nil ||
		instanceFlags.groupId.Get() == nil {
		return ERROR_USAGE
	}

	call := service.Clientupdate.List()
	call.DateStart(instanceFlags.start)
	call.DateEnd(instanceFlags.end)
	call.GroupId(instanceFlags.groupId.String())
	call.AppId(instanceFlags.appId.String())
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

func instanceListAppVersions(args []string, service *update.Service, out *tabwriter.Writer) int {
	if instanceFlags.appId.Get() == nil ||
		instanceFlags.groupId.Get() == nil {
		return ERROR_USAGE
	}

	call := service.Appversion.List()

	call.GroupId(instanceFlags.groupId.String())
	call.AppId(instanceFlags.appId.String())

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
