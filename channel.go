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

	cmdListChannels = &Command{
		Name:        "list-channels",
		Usage:       "[OPTION]...",
		Description: `List all channels for an application.`,
		Run:         listChannels,
	}
	cmdUpdateChannel = &Command{
		Name:        "update-channel",
		Usage:       "[OPTION]... <version>",
		Description: `Update a channel to a new version.`,
		Run:         updateChannel,
	}
)

func init() {
	cmdListChannels.Flags.Var(&channelFlags.appId, "app-id", "The application ID to list the channels of.")

	cmdUpdateChannel.Flags.Var(&channelFlags.appId, "app-id", "The application ID that the channel belongs to.")
	cmdUpdateChannel.Flags.Var(&channelFlags.channel, "channel", "The channel to update.")
}

func formatChannel(channel *update.AppChannel) string {
	return fmt.Sprintf("%s\t%s\n", channel.Label, channel.Version)
}

func listChannels(args []string, service *update.Service, out *tabwriter.Writer) int {
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

func updateChannel(args []string, service *update.Service, out *tabwriter.Writer) int {
	if channelFlags.appId.Get() == nil || channelFlags.channel.Get() == nil {
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
