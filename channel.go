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
			Name:        "update-channel",
			Usage:       "update-channel <appId> <channel> <version>",
			Description: `Update a given channel given a group, app, channel and version.`,
			Action:      handle(updateChannel),
		},
	}
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
