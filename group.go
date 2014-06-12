package main

import (
	"fmt"
	"log"
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
		autoPause      BoolFlag
		errorThreshold int64
		errorInterval  int64
		dropThreshold  int64
		dropInterval   int64
	}

	cmdGroup = &Command{
		Name:    "group",
		Summary: "Operations that manage groups in an applciation.",
		Subcommands: []*Command{
			cmdGroupList,
			cmdGroupCreate,
			cmdGroupDelete,
			cmdGroupUpdate,
			cmdGroupPause,
			cmdGroupUnpause,
			cmdGroupEvents,
			cmdGroupVersions,
		},
	}

	cmdGroupList = &Command{
		Name:    "group list",
		Usage:   "[OPTION]...",
		Summary: `List all of the groups that exist including their label, token and update state.`,
		Run:     groupList,
	}
	cmdGroupCreate = &Command{
		Name:    "group create",
		Usage:   "[OPTION]...",
		Summary: `Create a new group.`,
		Run:     groupCreate,
	}
	cmdGroupDelete = &Command{
		Name:    "group delete",
		Usage:   "[OPTION]...",
		Summary: `Delete a group.`,
		Run:     groupDelete,
	}
	cmdGroupUpdate = &Command{
		Name:        "group update",
		Usage:       "[OPTION]...",
		Description: `Update an existing group.`,
		Run:         groupUpdate,
	}
	cmdGroupPause = &Command{
		Name:    "group pause",
		Usage:   "[OPTION]...",
		Summary: `Pause a group's updates.`,
		Run:     groupPause,
	}
	cmdGroupUnpause = &Command{
		Name:    "group unpause",
		Usage:   "[OPTION]...",
		Summary: `Unpause a group's updates.`,
		Run:     groupUnpause,
	}
	cmdGroupVersions = &Command{
		Name:    "group versions",
		Usage:   "[OPTION]...",
		Summary: "List versions from clients by time.",
		Run:     groupVersions,
	}
	cmdGroupEvents = &Command{
		Name:    "group events",
		Usage:   "[OPTION]...",
		Summary: "List events from clients by time.",
		Run:     groupEvents,
	}
)

func init() {
	cmdGroupList.Flags.Var(&groupFlags.appId, "app-id",
		"Application containing the groups to list.")

	cmdGroupDelete.Flags.Var(&groupFlags.appId, "app-id",
		"Application with group to delete.")
	cmdGroupDelete.Flags.Var(&groupFlags.groupId, "group-id",
		"ID of group to delete.")

	cmdGroupCreate.Flags.Var(&groupFlags.appId, "app-id",
		"Application to add group to.")
	cmdGroupCreate.Flags.Var(&groupFlags.groupId, "group-id",
		"ID for the new group.")
	cmdGroupCreate.Flags.Var(&groupFlags.channel, "channel",
		"Channel to associate with the group.")
	cmdGroupCreate.Flags.Var(&groupFlags.label, "label",
		"Label describing the new group.")

	cmdGroupUpdate.Flags.Var(&groupFlags.appId, "app-id",
		"Application containing the group to update.")
	cmdGroupUpdate.Flags.Var(&groupFlags.groupId, "group-id",
		"ID for the group.")
	cmdGroupUpdate.Flags.Var(&groupFlags.label, "label",
		"Label describing the group")
	cmdGroupUpdate.Flags.Var(&groupFlags.channel, "channel",
		"Channel to associate with the group.")
	cmdGroupUpdate.Flags.Int64Var(&groupFlags.updateCount, "update-count",
		-1, "Number of instances per interval")
	cmdGroupUpdate.Flags.Int64Var(&groupFlags.updateInterval,
		"update-interval", -1, "Interval between updates")
	cmdGroupUpdate.Flags.Var(&groupFlags.autoPause,
		"auto-pause", "Enable/Disable AutoPause feature.")
	cmdGroupUpdate.Flags.Int64Var(&groupFlags.errorThreshold,
		"error-threshold", -1, "Errors per interval for autopause.")
	cmdGroupUpdate.Flags.Int64Var(&groupFlags.errorInterval,
		"error-interval", -1, "Interval for error threshold in seconds.")
	cmdGroupUpdate.Flags.Int64Var(&groupFlags.dropThreshold, "drop-threshold",
		-1, "Number of instances that can drop per drop interval.")
	cmdGroupUpdate.Flags.Int64Var(&groupFlags.dropInterval,
		"drop-interval", -1, "Interval for drop threshold in seconds.")

	cmdGroupPause.Flags.Var(&groupFlags.appId, "app-id",
		"Application containing the group to pause.")
	cmdGroupPause.Flags.Var(&groupFlags.groupId, "group-id",
		"ID for the group.")

	cmdGroupUnpause.Flags.Var(&groupFlags.appId, "app-id",
		"Application containing the group to unpause.")
	cmdGroupUnpause.Flags.Var(&groupFlags.groupId, "group-id",
		"ID for the group.")

	cmdGroupVersions.Flags.Var(&groupFlags.appId, "app-id",
		"Application containing the group.")
	cmdGroupVersions.Flags.Var(&groupFlags.groupId, "group-id",
		"ID for the group.")
	cmdGroupVersions.Flags.Int64Var(&groupFlags.resolution,
		"resolution", 60, "60, 3600 or 86400 seconds")
	cmdGroupVersions.Flags.Int64Var(&groupFlags.start, "start", 0,
		"Start date filter")
	cmdGroupVersions.Flags.Int64Var(&groupFlags.end, "end", 0,
		"End date filter")

	cmdGroupEvents.Flags.Var(&groupFlags.appId, "app-id",
		"Application containing the group.")
	cmdGroupEvents.Flags.Var(&groupFlags.groupId, "group-id",
		"ID for the group.")
	cmdGroupEvents.Flags.Int64Var(&groupFlags.resolution,
		"resolution", 60, "60, 3600 or 86400 seconds")
	cmdGroupEvents.Flags.Int64Var(&groupFlags.start, "start", 0,
		"Start date filter")
	cmdGroupEvents.Flags.Int64Var(&groupFlags.end, "end", 0,
		"End date filter")
}

const groupLegend = "Label\tApp\tChannel\tId\tUpdatesPaused\tAutoPause\tErrors\tDrops"

func formatGroup(group *update.Group) string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%t\t%t\t%d/%ds\t%d/%ds\n",
		group.Label, group.AppId, group.ChannelId, group.Id,
		group.UpdatesPaused, group.AutoPause, group.ErrorThreshold,
		group.ErrorInterval, group.DropThreshold, group.DropInterval)
}

func groupList(args []string, service *update.Service, out *tabwriter.Writer) int {
	if groupFlags.appId.Get() == nil {
		return ERROR_USAGE
	}

	listCall := service.Group.List(groupFlags.appId.String())
	list, err := listCall.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(out, groupLegend)
	for _, group := range list.Items {
		fmt.Fprintf(out, "%s", formatGroup(group))
	}

	out.Flush()
	return OK
}

func groupEvents(args []string, service *update.Service, out *tabwriter.Writer) int {
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

func groupVersions(args []string, service *update.Service, out *tabwriter.Writer) int {
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

func groupCreate(args []string, service *update.Service, out *tabwriter.Writer) int {
	if groupFlags.appId.Get() == nil ||
		groupFlags.groupId.Get() == nil ||
		groupFlags.channel.Get() == nil {
		return ERROR_USAGE
	}

	group := &update.Group{
		ChannelId: groupFlags.channel.String(),
		Id:        groupFlags.groupId.String(),
		Label:     groupFlags.label.String(),
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

func groupDelete(args []string, service *update.Service, out *tabwriter.Writer) int {
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

func groupPause(args []string, service *update.Service, out *tabwriter.Writer) int {
	return setUpdatesPaused(service, out, true)
}

func groupUnpause(args []string, service *update.Service, out *tabwriter.Writer) int {
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

func groupUpdate(args []string, service *update.Service, out *tabwriter.Writer) int {
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

	if autoPause := groupFlags.autoPause.Get(); autoPause != nil {
		group.AutoPause = *autoPause
	}
	if groupFlags.errorThreshold != -1 {
		group.ErrorThreshold = groupFlags.errorThreshold
	}
	if groupFlags.errorInterval != -1 {
		group.ErrorInterval = groupFlags.errorInterval
	}
	if groupFlags.dropThreshold != -1 {
		group.DropThreshold = groupFlags.dropThreshold
	}
	if groupFlags.dropInterval != -1 {
		group.DropInterval = groupFlags.dropInterval
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
