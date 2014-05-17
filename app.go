package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/coreos-inc/updatectl/client/update/v1"
	"github.com/coreos-inc/updatectl/third_party/github.com/codegangsta/cli"
)

func AppCommands() []cli.Command {
	return []cli.Command{
		{
			Name:  "list-apps",
			Usage: "list-apps",
			Description: `List all of the apps that exist including their label,
token and update state.`,
			Action: handle(listApps),
		},
		{
			Name:        "update-app",
			Usage:       "update-app <appId> <channelId> <label>",
			Description: `Update an app given a label.`,
			Action:      handle(updateApp),
		},
		{
			Name:        "delete-app",
			Usage:       "delete-app <appId>",
			Description: `Delete a app given an id.`,
			Action:      handle(deleteApp),
		},
	}
}

func formatApp(app *update.App) string {
	return fmt.Sprintf("%s\t%s\t%s\n", app.Id, app.Label, app.Description)
}

func listApps(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	listCall := service.App.List()
	list, err := listCall.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(out, "Id\tLabel\tDescription")
	for _, app := range list.Items {
		fmt.Fprintf(out, "%s", formatApp(app))
	}

	out.Flush()
}

func updateApp(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	args := c.Args()

	if len(args) != 3 {
		cli.ShowCommandHelp(c, "update-app")
		os.Exit(1)
	}
	appReq := &update.AppUpdateReq{Label: args[1], Description: args[2]}
	call := service.App.Update(args[0], appReq)
	app, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s", formatApp(app))

	out.Flush()
}

func deleteApp(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	args := c.Args()

	if len(args) != 1 {
		fmt.Println("App token is required")
		os.Exit(1)
	}

	call := service.App.Delete(args[0])
	app, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s", formatApp(app))

	out.Flush()
}
