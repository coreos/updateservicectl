package main

import (
	"fmt"
	"log"
	"text/tabwriter"

	"github.com/coreos-inc/updatectl/client/update/v1"
)

var (
	channelFlags struct {
		appId   StringFlag
		channel StringFlag
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
		Name:        "channel update",
		Usage:       "[OPTION]... <version>",
		Summary: `Update a channel for an application to a new version.`,
		Description: `Given an application ID (--app-id) and channel (--channel),
you can change the channel to a new version (<version>).`,
		Run:         channelUpdate,
	}
)

func init() {
	cmdChannelList.Flags.Var(&channelFlags.appId, "app-id", "The application ID to list the channels of.")

	cmdChannelUpdate.Flags.Var(&channelFlags.appId, "app-id", "The application ID that the channel belongs to.")
	cmdChannelUpdate.Flags.Var(&channelFlags.channel, "channel", "The channel to update.")
}

func formatChannel(channel *update.AppChannel) string {
	return fmt.Sprintf("%s\t%s\n", channel.Label, channel.Version)
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
	fmt.Fprint(out, "Label\tVersion\n")
	for _, channel := range list.Items {
		fmt.Fprintf(out, "%s", formatChannel(channel))
	}
	out.Flush()
	return OK
}

func channelUpdate(args []string, service *update.Service, out *tabwriter.Writer) int {
	if len(args) != 1 || channelFlags.appId.Get() == nil ||
		channelFlags.channel.Get() == nil{
		return ERROR_USAGE
	}

	channelReq := &update.ChannelRequest{Version: args[0]}

	call := service.Channel.Update(channelFlags.appId.String(), channelFlags.channel.String(), channelReq)
	channel, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s\n", channel.Version)
	out.Flush()
	return OK
}
