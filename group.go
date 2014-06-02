package main

import (
	"fmt"
	"log"
	"strconv"
	"text/tabwriter"

	"github.com/coreos-inc/updatectl/client/update/v1"
)

var (
	groupFlags struct {
		label          string
		channel        string
		start          int64
		end            int64
		resolution     int64
		updateCount    int64
		updateInterval int64
	}

	cmdListGroups = &Command{
		Name:    "list-groups",
		Usage:   "<appId>",
		Summary: `List all of the groups that exist including their label, token and update state.`,
		Run:     listGroups,
	}
	cmdNewGroup = &Command{
		Name:    "new-group",
		Usage:   "<appId> <channelId> <groupId> <appLabel>",
		Summary: `Create a new group given a label.`,
		Run:     newGroup,
	}
	cmdDeleteGroup = &Command{
		Name:    "delete-group",
		Usage:   "<appId> <groupId>",
		Summary: `Delete a group given a token.`,
		Run:     deleteGroup,
	}
	cmdUpdateGroup = &Command{
		Name:        "update-group",
		Usage:       "[OPTION]... <appId> <groupId>",
		Description: `Update an existing group.`,
		Run:         updateGroup,
	}
	cmdPauseGroup = &Command{
		Name:    "pause-group",
		Usage:   "<appId> <groupId>",
		Summary: `Pause a group given an id.`,
		Run:     pauseGroup,
	}
	cmdUnpauseGroup = &Command{
		Name:    "unpause-group",
		Usage:   "<appId> <groupId>",
		Summary: `Unpause a group given an id.`,
		Run:     unpauseGroup,
	}
	cmdRollupGroupVersions = &Command{
		Name:    "rollup-group-versions",
		Usage:   "[OPTION]... <appId> <groupId>",
		Summary: "Rollup versions from events by time.",
		Run:     rollupGroupVersions,
	}
	cmdRollupGroupEvents = &Command{
		Name:    "rollup-group-events",
		Usage:   "[OPTION]... <appId> <groupId>",
		Summary: "Rollup events from events by time.",
		Run:     rollupGroupEvents,
	}
)

func init() {
	cmdUpdateGroup.Flags.StringVar(&groupFlags.label, "label", "", "")
	cmdUpdateGroup.Flags.StringVar(&groupFlags.channel, "channel", "", "")
	cmdUpdateGroup.Flags.Int64Var(&groupFlags.updateCount, "updateCount", -1, "Number of instances per interval")
	cmdUpdateGroup.Flags.Int64Var(&groupFlags.updateInterval, "updateInterval", -1, "Interval between updates")

	cmdRollupGroupVersions.Flags.Int64Var(&groupFlags.resolution, "resolution", 60, "60, 3600 or 86400 seconds")
	cmdRollupGroupVersions.Flags.Int64Var(&groupFlags.start, "start", 0, "Start date filter")
	cmdRollupGroupVersions.Flags.Int64Var(&groupFlags.end, "end", 0, "End date filter")

	cmdRollupGroupEvents.Flags.Int64Var(&groupFlags.resolution, "resolution", 60, "60, 3600 or 86400 seconds")
	cmdRollupGroupEvents.Flags.Int64Var(&groupFlags.start, "start", 0, "Start date filter")
	cmdRollupGroupEvents.Flags.Int64Var(&groupFlags.end, "end", 0, "End date filter")
}

func formatGroup(group *update.Group) string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%v\t%v\n", group.Label, group.AppId, group.ChannelId,
		group.Id, strconv.FormatBool(group.UpdatesPaused), group.UpdateCount, group.UpdateInterval)
}

func listGroups(args []string, service *update.Service, out *tabwriter.Writer) int {
	if len(args) != 1 {
		return ERROR_USAGE
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
	return OK
}

func rollupGroupEvents(args []string, service *update.Service, out *tabwriter.Writer) int {
	if len(args) != 2 {
		return ERROR_USAGE
	}

	call := service.Group.Requests.Events.Rollup(args[0], args[1], groupFlags.start, groupFlags.end)
	call.Resolution(groupFlags.resolution)
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
	return OK
}

func rollupGroupVersions(args []string, service *update.Service, out *tabwriter.Writer) int {
	if len(args) != 2 {
		return ERROR_USAGE
	}

	call := service.Group.Requests.Versions.Rollup(args[0], args[1], groupFlags.start, groupFlags.end)
	call.Resolution(groupFlags.resolution)
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
	return OK
}

func newGroup(args []string, service *update.Service, out *tabwriter.Writer) int {
	if len(args) != 4 {
		return ERROR_USAGE
	}
	group := &update.Group{ChannelId: args[1], Id: args[2], Label: args[3]}
	call := service.Group.Insert(args[0], group)
	group, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s", formatGroup(group))

	out.Flush()
	return OK
}

func deleteGroup(args []string, service *update.Service, out *tabwriter.Writer) int {
	if len(args) != 2 {
		return ERROR_USAGE
	}

	call := service.Group.Delete(args[0], args[1])
	group, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s", formatGroup(group))

	out.Flush()
	return OK
}

func pauseGroup(args []string, service *update.Service, out *tabwriter.Writer) int {
	if len(args) != 2 {
		return ERROR_USAGE
	}
	return setUpdatesPaused(args, service, out, true)
}

func unpauseGroup(args []string, service *update.Service, out *tabwriter.Writer) int {
	if len(args) != 2 {
		return ERROR_USAGE
	}
	return setUpdatesPaused(args, service, out, false)
}

// Helper function for pause/unpause-group commands
func setUpdatesPaused(args []string, service *update.Service, out *tabwriter.Writer, paused bool) int {
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
	return OK
}

func updateGroup(args []string, service *update.Service, out *tabwriter.Writer) int {
	if len(args) != 2 {
		return ERROR_USAGE
	}

	call := service.Group.Get(args[0], args[1])
	group, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	checkUpdatePooling := false
	if groupFlags.updateCount != -1 {
		group.UpdateCount = groupFlags.updateCount
		checkUpdatePooling = true
	}
	if groupFlags.updateInterval != -1 {
		group.UpdateInterval = groupFlags.updateInterval
		checkUpdatePooling = true
	}
	if groupFlags.label != "" {
		group.Label = groupFlags.label
	}
	if groupFlags.channel != "" {
		group.ChannelId = groupFlags.channel
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
	return OK
}
