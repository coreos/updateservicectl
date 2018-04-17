package main

import (
	"fmt"
	"log"
	"strconv"
	"text/tabwriter"

	"github.com/coreos/updateservicectl/client/update/v1"
)

var (
	groupFlags struct {
		label         StringFlag
		channel       StringFlag
		appId         StringFlag
		groupId       StringFlag
		oemBlacklist  StringFlag
		start         int64
		end           int64
		resolution    int64
		updatePercent float64
	}

	cmdGroup = &Command{
		Name:    "group",
		Summary: "Operations that manage groups in an application.",
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
	cmdGroupUpdate.Flags.Var(&groupFlags.oemBlacklist, "oem-blacklist",
		"Comma-separated list of OEMs to exclude from updates.")
	cmdGroupUpdate.Flags.Float64Var(&groupFlags.updatePercent,
		"update-percent", -1, "Percentage of machines to update")

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

const groupHeader = "Label\tApp\tChannel\tId\tPaused\tPercent\n"

func formatGroup(group *update.Group) string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%v\n",
		group.Label, group.AppId, group.ChannelId,
		group.Id, strconv.FormatBool(group.UpdatesPaused),
		group.UpdatePercent)
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

	fmt.Fprint(out, groupHeader)
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

	fmt.Fprint(out, groupHeader)
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

	fmt.Fprint(out, groupHeader)
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

	fmt.Fprint(out, groupHeader)
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

	if groupFlags.updatePercent != -1 {
		group.UpdatePercent = groupFlags.updatePercent
	}
	if groupFlags.label.Get() != nil {
		group.Label = groupFlags.label.String()
	}
	if groupFlags.channel.Get() != nil {
		group.ChannelId = groupFlags.channel.String()
	}
	if groupFlags.oemBlacklist.Get() != nil {
		group.OemBlacklist = groupFlags.oemBlacklist.String()
	}

	updateCall := service.Group.Patch(groupFlags.appId.String(), groupFlags.groupId.String(), group)
	group, err = updateCall.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(out, groupHeader)
	fmt.Fprintf(out, "%s", formatGroup(group))

	out.Flush()
	return OK
}
