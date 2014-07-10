package main

import (
	"fmt"
	"log"
	"text/tabwriter"

	"github.com/coreos/updatectl/client/update/v1"
)

var (
	channelFlags struct {
		appId   StringFlag
		channel StringFlag
		version StringFlag
		publish bool
	}

	cmdChannel = &Command{
		Name:    "channel",
		Summary: "Manage channels for an application",
		Subcommands: []*Command{
			cmdChannelList,
			cmdChannelUpdate,
		},
	}

	cmdChannelList = &Command{
		Name:        "channel list",
		Usage:       "[OPTION]...",
		Description: `List all channels for an application.`,
		Run:         channelList,
	}
	cmdChannelUpdate = &Command{
		Name:    "channel update",
		Usage:   "[OPTION]...",
		Summary: `Update the version and publish state for an application channel.`,
		Description: `Given an application ID (--app-id) and channel (--channel),
you can change the channel to a new version (--version), or set the publish state (--publish).`,
		Run: channelUpdate,
	}
)

func init() {
	cmdChannelList.Flags.Var(&channelFlags.appId, "app-id", "The application ID to list the channels of.")

	cmdChannelUpdate.Flags.Var(&channelFlags.appId, "app-id", "The application ID that the channel belongs to.")
	cmdChannelUpdate.Flags.Var(&channelFlags.channel, "channel", "The channel to update.")
	cmdChannelUpdate.Flags.BoolVar(&channelFlags.publish, "publish", false, "Publish or unpublish the channel.")
	cmdChannelUpdate.Flags.Var(&channelFlags.version, "version", "The version to update the channel to.")
}

func formatChannel(channel *update.AppChannel) string {
	return fmt.Sprintf("%s\t%s\t%t\t%s\n", channel.Label, channel.Version, channel.Publish, channel.Upstream)
}

func channelList(args []string, service *update.Service, out *tabwriter.Writer) int {
	if channelFlags.appId.Get() == nil {
		return ERROR_USAGE
	}

	listCall := service.Channel.List(channelFlags.appId.String())
	list, err := listCall.Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(out, "Label\tVersion\tPublish\tUpstream\n")
	for _, channel := range list.Items {
		fmt.Fprintf(out, "%s", formatChannel(channel))
	}
	out.Flush()
	return OK
}

func channelUpdate(args []string, service *update.Service, out *tabwriter.Writer) int {
	if channelFlags.version.Get() == nil || channelFlags.appId.Get() == nil || channelFlags.channel.Get() == nil {
		return ERROR_USAGE
	}

	channelReq := &update.ChannelRequest{Version: *channelFlags.version.Get(), Publish: channelFlags.publish}

	call := service.Channel.Update(channelFlags.appId.String(), channelFlags.channel.String(), channelReq)
	channel, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s", formatChannel(channel))
	out.Flush()
	return OK
}
