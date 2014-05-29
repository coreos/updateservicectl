package main

import (
	"log"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/coreos-inc/updatectl/auth"
	"github.com/coreos-inc/updatectl/version"
	"github.com/coreos-inc/updatectl/client/update/v1"
	"github.com/coreos-inc/updatectl/third_party/github.com/codegangsta/cli"
)

type handlerFunc func(*cli.Context, *update.Service, *tabwriter.Writer)

func getHawkClient(c *cli.Context) *http.Client {
	return &http.Client{Transport: &auth.HawkRoundTripper{
		c.GlobalString("user"),
		c.GlobalString("key"),
	}}
}

func handle(fn handlerFunc) func(c *cli.Context) {
	out := new(tabwriter.Writer)
	out.Init(os.Stdout, 0, 8, 1, '\t', 0)

	return func(c *cli.Context) {
		client := getHawkClient(c)

		service, err := update.New(client)
		if err != nil {
			log.Fatal(err)
		}

		service.BasePath = c.GlobalString("server") + "/_ah/api/update/v1/"
		fn(c, service, out)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "updatectl"
	app.Usage = "updatectl is a command line driven interface to the roller."
	app.Action = handle(listGroups)
	app.Version = version.SemVersion.String()
	app.Flags = []cli.Flag{
		cli.StringFlag{"server, s", "http://localhost:8000", "Update server to connect to"},
		cli.BoolFlag{"debug, D", "Output debugging info to stderr"},
		cli.StringFlag{"user, u", os.Getenv("UPDATECTL_USER"), "API Username"},
		cli.StringFlag{"key, k", os.Getenv("UPDATECTL_KEY"), "API Key"},
	}

	app.Commands = append(app.Commands, GroupCommands()...)
	app.Commands = append(app.Commands, AppCommands()...)
	app.Commands = append(app.Commands, ChannelCommands()...)
	app.Commands = append(app.Commands, PackageCommands()...)
	app.Commands = append(app.Commands, WatchCommands()...)
	app.Commands = append(app.Commands, AdminCommands()...)
	app.Commands = append(app.Commands, ClientUpdateCommands()...)
	app.Commands = append(app.Commands, FakeClientsCommand()...)
	app.Run(os.Args)
}
