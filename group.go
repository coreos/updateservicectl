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
		label          StringFlag
		channel        StringFlag
		appId          StringFlag
		groupId        StringFlag
		start          int64
		end            int64
		resolution     int64
		updateCount    int64
		updateInterval int64
	}

	cmdListGroups = &Command{
		Name:    "list-groups",
		Usage:   "[OPTION]...",
		Summary: `List all of the groups that exist including their label, token and update state.`,
		Run:     listGroups,
	}
	cmdNewGroup = &Command{
		Name:    "new-group",
		Usage:   "[OPTION]...",
		Summary: `Create a new group.`,
		Run:     newGroup,
	}
	cmdDeleteGroup = &Command{
		Name:    "delete-group",
		Usage:   "[OPTION]...",
		Summary: `Delete a group.`,
		Run:     deleteGroup,
	}
	cmdUpdateGroup = &Command{
		Name:        "update-group",
		Usage:       "[OPTION]...",
		Description: `Update an existing group.`,
		Run:         updateGroup,
	}
	cmdPauseGroup = &Command{
		Name:    "pause-group",
		Usage:   "[OPTION]...",
		Summary: `Pause a group's updates.`,
		Run:     pauseGroup,
	}
	cmdUnpauseGroup = &Command{
		Name:    "unpause-group",
		Usage:   "[OPTION]...",
		Summary: `Unpause a group's updates.`,
		Run:     unpauseGroup,
	}
	cmdRollupGroupVersions = &Command{
		Name:    "rollup-group-versions",
		Usage:   "[OPTION]...",
		Summary: "Rollup versions from events by time.",
		Run:     rollupGroupVersions,
	}
	cmdRollupGroupEvents = &Command{
		Name:    "rollup-group-events",
		Usage:   "[OPTION]...",
		Summary: "Rollup events from events by time.",
		Run:     rollupGroupEvents,
	}
)

func init() {
	cmdListGroups.Flags.Var(&groupFlags.appId, "app-id",
		"Application containing the groups to list.")

	cmdDeleteGroup.Flags.Var(&groupFlags.appId, "app-id",
		"Application with group to delete.")
	cmdDeleteGroup.Flags.Var(&groupFlags.groupId, "group-id",
		"ID of group to delete.")


	cmdNewGroup.Flags.Var(&groupFlags.appId, "app-id",
		"Application to add group to.")
	cmdNewGroup.Flags.Var(&groupFlags.groupId, "group-id",
		 "ID for the new group.")
	cmdNewGroup.Flags.Var(&groupFlags.channel, "channel",
		 "Channel to associate with the group.")
	cmdNewGroup.Flags.Var(&groupFlags.label, "label",
		"Label describing the new group.")

	cmdUpdateGroup.Flags.Var(&groupFlags.appId, "app-id",
		 "Application containing the group to update.")
	cmdUpdateGroup.Flags.Var(&groupFlags.groupId, "group-id",
		 "ID for the group.")
	cmdUpdateGroup.Flags.Var(&groupFlags.label, "label",
		"Label describing the group")
	cmdUpdateGroup.Flags.Var(&groupFlags.channel, "channel",
		"Channel to associate with the group.")
	cmdUpdateGroup.Flags.Int64Var(&groupFlags.updateCount, "update-count",
		-1, "Number of instances per interval")
	cmdUpdateGroup.Flags.Int64Var(&groupFlags.updateInterval,
		"update-interval", -1, "Interval between updates")

	cmdPauseGroup.Flags.Var(&groupFlags.appId, "app-id",
		 "Application containing the group to pause.")
	cmdPauseGroup.Flags.Var(&groupFlags.groupId, "group-id",
		 "ID for the group.")

	cmdUnpauseGroup.Flags.Var(&groupFlags.appId, "app-id",
		 "Application containing the group to unpause.")
	cmdUnpauseGroup.Flags.Var(&groupFlags.groupId, "group-id",
		 "ID for the group.")

	cmdRollupGroupVersions.Flags.Var(&groupFlags.appId, "app-id",
		 "Application containing the group.")
	cmdRollupGroupVersions.Flags.Var(&groupFlags.groupId, "group-id",
		 "ID for the group.")
	cmdRollupGroupVersions.Flags.Int64Var(&groupFlags.resolution,
		"resolution", 60, "60, 3600 or 86400 seconds")
	cmdRollupGroupVersions.Flags.Int64Var(&groupFlags.start, "start", 0,
		"Start date filter")
	cmdRollupGroupVersions.Flags.Int64Var(&groupFlags.end, "end", 0,
		"End date filter")

	cmdRollupGroupEvents.Flags.Var(&groupFlags.appId, "app-id",
		 "Application containing the group.")
	cmdRollupGroupEvents.Flags.Var(&groupFlags.groupId, "group-id",
		 "ID for the group.")
	cmdRollupGroupEvents.Flags.Int64Var(&groupFlags.resolution,
		"resolution", 60, "60, 3600 or 86400 seconds")
	cmdRollupGroupEvents.Flags.Int64Var(&groupFlags.start, "start", 0,
		"Start date filter")
	cmdRollupGroupEvents.Flags.Int64Var(&groupFlags.end, "end", 0,
		"End date filter")
}

func formatGroup(group *update.Group) string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%v\t%v\n", group.Label, group.AppId, group.ChannelId,
		group.Id, strconv.FormatBool(group.UpdatesPaused), group.UpdateCount, group.UpdateInterval)
}

func listGroups(args []string, service *update.Service, out *tabwriter.Writer) int {
	if groupFlags.appId.Get() == nil {
		return ERROR_USAGE
	}

	listCall := service.Group.List(groupFlags.appId.String())
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
	if groupFlags.appId.Get() == nil || groupFlags.groupId.Get() == nil {
		return ERROR_USAGE
	}

	call := service.Group.Requests.Events.Rollup(
		groupFlags.appId.String(),
		groupFlags.groupId.String(),
		groupFlags.start,
		groupFlags.end,
	)
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
	if groupFlags.appId.Get() == nil || groupFlags.groupId.Get() == nil {
		return ERROR_USAGE
	}

	call := service.Group.Requests.Versions.Rollup(
		groupFlags.appId.String(),
		groupFlags.groupId.String(),
		groupFlags.start,
		groupFlags.end,
	)
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
	if groupFlags.appId.Get() == nil ||
		groupFlags.groupId.Get() == nil ||
		groupFlags.channel.Get() == nil {
		return ERROR_USAGE
	}

	group := &update.Group{
		ChannelId: groupFlags.channel.String(),
		Id: groupFlags.groupId.String(),
		Label: groupFlags.label.String(),
	}
	call := service.Group.Insert(groupFlags.appId.String(), group)
	group, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s", formatGroup(group))

	out.Flush()
	return OK
}

func deleteGroup(args []string, service *update.Service, out *tabwriter.Writer) int {
	if groupFlags.appId.Get() == nil ||
		groupFlags.groupId.Get() == nil {
		return ERROR_USAGE
	}

	call := service.Group.Delete(groupFlags.appId.String(), groupFlags.groupId.String())
	group, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s", formatGroup(group))

	out.Flush()
	return OK
}

func pauseGroup(args []string, service *update.Service, out *tabwriter.Writer) int {
	return setUpdatesPaused(service, out, true)
}

func unpauseGroup(args []string, service *update.Service, out *tabwriter.Writer) int {
	return setUpdatesPaused(service, out, false)
}

// Helper function for pause/unpause-group commands
func setUpdatesPaused(service *update.Service, out *tabwriter.Writer, paused bool) int {
	if groupFlags.appId.Get() == nil ||
		groupFlags.groupId.Get() == nil {
		return ERROR_USAGE
	}

	call := service.Group.Get(groupFlags.appId.String(), groupFlags.groupId.String())
	group, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	group.UpdatesPaused = paused

	updateCall := service.Group.Patch(groupFlags.appId.String(), groupFlags.groupId.String(), group)
	group, err = updateCall.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s", formatGroup(group))

	out.Flush()
	return OK
}

func updateGroup(args []string, service *update.Service, out *tabwriter.Writer) int {
	if groupFlags.appId.Get() == nil ||
		groupFlags.groupId.Get() == nil {
		return ERROR_USAGE
	}

	call := service.Group.Get(groupFlags.appId.String(), groupFlags.groupId.String())
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
	if groupFlags.label.Get() != nil {
		group.Label = groupFlags.label.String()
	}
	if groupFlags.channel.Get() != nil {
		group.ChannelId = groupFlags.channel.String()
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

	updateCall := service.Group.Patch(groupFlags.appId.String(), groupFlags.groupId.String(), group)
	group, err = updateCall.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s", formatGroup(group))

	out.Flush()
	return OK
}
