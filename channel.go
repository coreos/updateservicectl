package main

import (
	"fmt"
	"log"
	"text/tabwriter"

	"github.com/coreos-inc/updatectl/client/update/v1"
)

var (
	channelFlags struct {
		appId   string
		channel string
	}

	cmdListChannels = &Command{
		Name:        "list-channels",
		Usage:       "<appId>",
		Description: `List all channels for an application.`,
		Run:         listChannels,
	}
	cmdUpdateChannel = &Command{
		Name:        "update-channel",
		Usage:       "[OPTIONS]... <version>",
		Description: `Update a channel to a new version.`,
		Run:         updateChannel,
	}
)

func init() {
	cmdUpdateChannel.Flags.StringVar(&channelFlags.appId, "appid", REQUIRED_FLAG, "The application ID that the channel belongs to.")
	cmdUpdateChannel.Flags.StringVar(&channelFlags.channel, "channel", REQUIRED_FLAG, "The channel to update.")
}

func formatChannel(channel *update.AppChannel) string {
	return fmt.Sprintf("%s\t%s\n", channel.Label, channel.Version)
}

func listChannels(args []string, service *update.Service, out *tabwriter.Writer) int {
	if len(args) != 1 {
		return ERROR_USAGE
	}

	listCall := service.Channel.List(args[0])
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
	if len(args) != 1 {
		return ERROR_USAGE
	}

	if channelFlags.appId == REQUIRED_FLAG || channelFlags.channel == REQUIRED_FLAG {
		return ERROR_USAGE
	}

	channelReq := &update.ChannelRequest{Version: args[0]}

	call := service.Channel.Update(channelFlags.appId, channelFlags.channel, channelReq)
	channel, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s\n", channel.Version)
	out.Flush()
	return OK
}
