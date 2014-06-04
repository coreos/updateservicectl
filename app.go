package main

import (
	"fmt"
	"log"
	"text/tabwriter"

	"code.google.com/p/go-uuid/uuid"
	"github.com/coreos-inc/updatectl/client/update/v1"
)

var (
	cmdListApps = &Command{
		Name:  "list-apps",
		Description: `List all of the apps that exist including their label,
token and update state.`,
		Run: listApps,
	}
	cmdCreateApp = &Command{
		Name:        "create-app",
		Usage:       "[<appId>] <label> <description>",
		Description: `Create a new application. If appId is not provided it will be created randomly.`,
		Run:         createApp,
	}
	cmdUpdateApp = &Command{
		Name:        "update-app",
		Usage:       "<appId> <label> <description>",
		Description: `Update an app given a label.`,
		Run:         updateApp,
	}
	cmdDeleteApp = &Command{
		Name:        "delete-app",
		Usage:       "<appId>",
		Description: `Delete an app given an id.`,
		Run:         deleteApp,
	}
)

func formatApp(app *update.App) string {
	return fmt.Sprintf("%s\t%s\t%s\n", app.Id, app.Label, app.Description)
}

func listApps(args []string, service *update.Service, out *tabwriter.Writer) int {
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
	return OK
}

func createApp(args []string, service *update.Service, out *tabwriter.Writer) int {
	if !(len(args) == 2 || len(args) == 3) {
		return ERROR_USAGE
	}

	if len(args) == 2 {
		appId := uuid.New()
		args = append([]string{appId}, args...)
	}

	updateAppHelper(args, service, out)
	return OK
}

func updateApp(args []string, service *update.Service, out *tabwriter.Writer) int {
	if len(args) != 3 {
		return ERROR_USAGE
	}
	return updateAppHelper(args, service, out)
}

func updateAppHelper(args []string, service *update.Service, out *tabwriter.Writer) int {
	appReq := &update.AppUpdateReq{Label: args[1], Description: args[2]}
	call := service.App.Update(args[0], appReq)
	app, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s", formatApp(app))

	out.Flush()
	return OK
}

func deleteApp(args []string, service *update.Service, out *tabwriter.Writer) int {
	if len(args) != 1 {
		return ERROR_USAGE
	}

	call := service.App.Delete(args[0])
	app, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s", formatApp(app))

	out.Flush()
	return OK
}
