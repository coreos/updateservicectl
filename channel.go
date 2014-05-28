package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/coreos-inc/updatectl/client/update/v1"
	"github.com/coreos-inc/updatectl/third_party/github.com/codegangsta/cli"
)

func ChannelCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "list-channels",
			Usage:       "list-channels <appId>",
			Description: `List all channels for an application.`,
			Action:      handle(listChannels),
		},
		{
			Name:        "update-channel",
			Usage:       "update-channel <appId> <channel> <version>",
			Description: `Update a given channel given a group, app, channel and version.`,
			Action:      handle(updateChannel),
		},
	}
}

func formatChannel(channel *update.AppChannel) string {
	return fmt.Sprintf("%s\t%s\n", channel.Label, channel.Version)
}

func listChannels(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	args := c.Args()

	if len(args) != 1 {
		cli.ShowCommandHelp(c, "list-channels")
		os.Exit(1)
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
}


func updateChannel(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	args := c.Args()

	if len(args) != 3 {
		fmt.Println("usage: <appid> <channel> <version>")
		os.Exit(1)
	}

	channelReq := &update.ChannelRequest{Version: args[2]}

	call := service.Channel.Update(args[0], args[1], channelReq)
	channel, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s\n", channel.Version)
	out.Flush()
}
