package main

import (
	"fmt"
	"log"
	"text/tabwriter"

	update "github.com/coreos-inc/updatectl/client/update/v1"
	"github.com/codegangsta/cli"
)

func ClientUpdateCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "list-clientupdates",
			Usage:       "list-clientupdates [OPTION]...",
			Description: "Generates a list of client updates.",
			Action:      handle(listClientUpdates),
			Flags: []cli.Flag{
				cli.StringFlag{"group-id", "", "Group id"},
				cli.StringFlag{"app-id", "", "App id"},
				cli.IntFlag{"start", 0, "Start date filter"},
				cli.IntFlag{"end", 0, "End date filter"},
			},
		},
		{
			Name:        "list-appversions",
			Usage:       "list-appversions [OPTION]...",
			Description: "Generates a list of Apps/versions with client count.",
			Action:      handle(listAppVersions),
			Flags: []cli.Flag{
				cli.IntFlag{"start", 0, "Start date filter"},
				cli.IntFlag{"end", 0, "End date filter"},
			},
		},
	}
}

func listClientUpdates(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	call := service.Clientupdate.List()
	call.DateStart(int64(c.Int("start")))
	call.DateEnd(int64(c.Int("end")))
	call.GroupId(c.String("group-id"))
	call.AppId(c.String("app-id"))
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
}

func listAppVersions(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	call := service.Appversion.List()

	call.GroupId(c.String("group-id"))
	call.AppId(c.String("app-id"))

	if c.Int("start") != 0 {
		call.DateStart(int64(c.Int("start")))
	}

	if c.Int("end") != 0 {
		call.DateEnd(int64(c.Int("end")))
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
}
