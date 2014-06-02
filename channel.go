package main

import (
	"fmt"
	"log"
	"text/tabwriter"

	"github.com/coreos-inc/updatectl/client/update/v1"
)

var (
	cmdListChannels = &Command{
		Name:        "list-channels",
		Usage:       "<appId>",
		Description: `List all channels for an application.`,
		Run:         listChannels,
	}
	cmdUpdateChannel = &Command{
		Name:        "update-channel",
		Usage:       "<appId> <channel> <version>",
		Description: `Update a given channel given a group, app, channel and version.`,
		Run:         updateChannel,
	}
)

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
	if len(args) != 3 {
		return ERROR_USAGE
	}

	channelReq := &update.ChannelRequest{Version: args[2]}

	call := service.Channel.Update(args[0], args[1], channelReq)
	channel, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s\n", channel.Version)
	out.Flush()
	return OK
}
