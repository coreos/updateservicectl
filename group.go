package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/coreos-inc/updatectl/client/update/v1"
	"github.com/codegangsta/cli"
)

func GroupCommands() []cli.Command {
	return []cli.Command{
		{
			Name:  "list-groups",
			Usage: "list-groups <appId>",
			Description: `List all of the groups that exist including their label,
token and update state.`,
			Action: handle(listGroups),
		},
		{
			Name:        "new-group",
			Usage:       "new-group <appId> <channelId> <groupId> <appLabel>",
			Description: `Create a new group given a label.`,
			Action:      handle(newGroup),
		},
		{
			Name:        "update-group",
			Usage:       "update-group [OPTION]... <appId> <groupId>",
			Description: `Update an existing group.`,
			Action:      handle(updateGroup),
			Flags: []cli.Flag{
				cli.StringFlag{"label", "", "Label to easily identify this group"},
				cli.StringFlag{"channel", "", "Name of channel (defaults to nil)"},
				cli.IntFlag{"updateCount", 0, "Number of instances per interval"},
				cli.IntFlag{"updateInterval", 0, "Interval between updates"},
			},
		},
		{
			Name:        "delete-group",
			Usage:       "delete-group <appId> <groupId>",
			Description: `Delete a group given a token.`,
			Action:      handle(deleteGroup),
		},
		{
			Name:        "pause-group",
			Usage:       "pause-group <appId> <groupId>",
			Description: `Pause a group given an id.`,
			Action:      handle(pauseGroup),
		},
		{
			Name:        "unpause-group",
			Usage:       "unpause-group <appId> <groupId>",
			Description: `Unpause a group given an id.`,
			Action:      handle(unpauseGroup),
		},
		{
			Name:        "rollup-group-versions",
			Usage:       "rollup-group-versions [OPTION]... <appId> <groupId>",
			Description: "Rollup versions from events by time.",
			Action:      handle(rollupGroupVersions),
			Flags: []cli.Flag{
				cli.IntFlag{"resolution", 60, "60, 3600 or 86400 seconds"},
				cli.IntFlag{"start", 0, "Start date filter"},
				cli.IntFlag{"end", 0, "End date filter"},
			},
		},
		{
			Name:        "rollup-group-events",
			Usage:       "rollup-group-events [OPTION]... <appId> <groupId>",
			Description: "Rollup versions from events by time.",
			Action:      handle(rollupGroupEvents),
			Flags: []cli.Flag{
				cli.IntFlag{"resolution", 60, "60, 3600 or 86400 seconds"},
				cli.IntFlag{"start", 0, "Start date filter"},
				cli.IntFlag{"end", 0, "End date filter"},
			},
		},
	}
}

func formatGroup(group *update.Group) string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%v\t%v\n", group.Label, group.AppId, group.ChannelId,
		group.Id, strconv.FormatBool(group.UpdatesPaused), group.UpdateCount, group.UpdateInterval)
}

func listGroups(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	args := c.Args()

	if len(args) != 1 {
		fmt.Println("AppId required")
		os.Exit(1)
	}

	listCall := service.Group.List(args[0])
	list, err := listCall.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(out, "Label\tApp\tChannel\tId\tUpdatesPaused")
	for _, group := range list.Items {
		fmt.Fprintf(out, "%s", formatGroup(group))
	}

	out.Flush()
}

func rollupGroupEvents(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	args := c.Args()

	if len(args) != 2 {
		cli.ShowCommandHelp(c, "rollup-group-events")
		os.Exit(1)
	}

	call := service.Group.Requests.Events.Rollup(args[0], args[1], int64(c.Int("start")), int64(c.Int("end")))
	call.Resolution(int64(c.Int("resolution")))
	list, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(out, "Version\tType\tResult\tTimestamp\tCount")
	for _, i := range list.Items {
		for _, j := range i.Values {
			fmt.Fprintf(out, "%s\t%s\t%s\t%d\t%d\n",
				i.Version, i.Type, i.Result, j.Timestamp, j.Count)
		}
	}
	out.Flush()
}

func rollupGroupVersions(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	args := c.Args()

	if len(args) != 2 {
		cli.ShowCommandHelp(c, "rollup-group-versions")
		os.Exit(1)
	}

	call := service.Group.Requests.Versions.Rollup(args[0], args[1], int64(c.Int("start")), int64(c.Int("end")))
	call.Resolution(int64(c.Int("resolution")))
	list, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(out, "Version\tTimestamp\tCount")
	for _, i := range list.Items {
		for _, j := range i.Values {
			fmt.Fprintf(out, "%s\t%d\t%d\n",
				i.Version, j.Timestamp, j.Count)
		}
	}
	out.Flush()
}

func newGroup(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	args := c.Args()

	if len(args) != 4 {
		fmt.Println("AppId, ChannelId, GroupId and app label required")
		os.Exit(1)
	}
	group := &update.Group{ChannelId: args[1], Id: args[2], Label: args[3]}
	call := service.Group.Insert(args[0], group)
	group, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s", formatGroup(group))

	out.Flush()
}

func deleteGroup(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	args := c.Args()

	if len(args) != 2 {
		fmt.Println("AppId and GroupId is required")
		os.Exit(1)
	}

	call := service.Group.Delete(args[0], args[1])
	group, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s", formatGroup(group))

	out.Flush()
}

func pauseGroup(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	setUpdatesPaused(c, service, out, true)
}

func unpauseGroup(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	setUpdatesPaused(c, service, out, false)
}

// Helper function for pause/unpause-group commands
func setUpdatesPaused(c *cli.Context, service *update.Service, out *tabwriter.Writer, paused bool) {
	args := c.Args()

	if len(args) != 2 {
		fmt.Println("AppId and GroupId is required")
		os.Exit(1)
	}

	call := service.Group.Get(args[0], args[1])
	group, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	group.UpdatesPaused = paused

	updateCall := service.Group.Patch(args[0], args[1], group)
	group, err = updateCall.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s", formatGroup(group))

	out.Flush()
}

func updateGroup(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	args := c.Args()

	if len(args) != 2 {
		cli.ShowCommandHelp(c, "update-group")
		os.Exit(1)
	}

	call := service.Group.Get(args[0], args[1])
	group, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	checkUpdatePooling := false
	if c.IsSet("updateCount") {
		group.UpdateCount = int64(c.Int("updateCount"))
		checkUpdatePooling = true
	}
	if c.IsSet("updateInterval") {
		group.UpdateInterval = int64(c.Int("updateInterval"))
		checkUpdatePooling = true
	}
	if c.IsSet("label") {
		group.Label = c.String("label")
	}
	if c.IsSet("channel") {
		group.ChannelId = c.String("channel")
	}

	// set update pooling based on other flags
	// this only changes if the user changed a value
	if checkUpdatePooling {
		if group.UpdateCount == 0 && group.UpdateInterval == 0 {
			group.UpdatePooling = false
		} else {
			group.UpdatePooling = true
		}
	}

	updateCall := service.Group.Patch(args[0], args[1], group)
	group, err = updateCall.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s", formatGroup(group))

	out.Flush()
}
