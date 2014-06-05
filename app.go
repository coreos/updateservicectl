package main

import (
	"fmt"
	"log"
	"text/tabwriter"

	"code.google.com/p/go-uuid/uuid"
	"github.com/coreos-inc/updatectl/client/update/v1"
)

var (
	appFlags struct {
		appId       StringFlag
		label       StringFlag
		description StringFlag
	}

	cmdListApps = &Command{
		Name:        "list-apps",
		Description: `List all of the apps that exist including their label, token and update state.`,
		Run:         listApps,
	}
	cmdCreateApp = &Command{
		Name:        "create-app",
		Usage:       "[OPTION]...",
		Description: `Create a new application.`,
		Run:         createApp,
	}
	cmdUpdateApp = &Command{
		Name:        "update-app",
		Usage:       "[OPTION]...",
		Description: `Update an app's label or description.`,
		Run:         updateApp,
	}
	cmdDeleteApp = &Command{
		Name:        "delete-app",
		Usage:       "[OPTION]...",
		Description: `Delete an app.`,
		Run:         deleteApp,
	}
)

func init() {
	cmdCreateApp.Flags.Var(&appFlags.appId, "app-id", "Application UUID. If not provided, one will be randomly generated.")
	cmdCreateApp.Flags.Var(&appFlags.label, "label", "New application label.")
	cmdCreateApp.Flags.Var(&appFlags.description, "description", "New application description.")

	cmdUpdateApp.Flags.Var(&appFlags.appId, "app-id",  "Application ID to update.")
	cmdUpdateApp.Flags.Var(&appFlags.label, "label", "Set application label.")
	cmdUpdateApp.Flags.Var(&appFlags.description, "description", "Set application description.")

	cmdDeleteApp.Flags.Var(&appFlags.appId, "app-id", "Application ID to delete.")
}

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
	if appFlags.appId.Get() == nil  {
		appFlags.appId.Set(uuid.New())
	}
	return updateAppHelper(service, out)
}

func updateApp(args []string, service *update.Service, out *tabwriter.Writer) int {
	if appFlags.appId.Get() == nil || appFlags.label.Get() == nil || appFlags.description.Get() == nil {
		return ERROR_USAGE
	}

	return updateAppHelper(service, out)
}

func updateAppHelper(service *update.Service, out *tabwriter.Writer) int {
	appReq := &update.AppUpdateReq{Label: appFlags.label.String(), Description: appFlags.description.String()}
	call := service.App.Update(appFlags.appId.String(), appReq)
	app, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s", formatApp(app))

	out.Flush()
	return OK
}

func deleteApp(args []string, service *update.Service, out *tabwriter.Writer) int {
	if appFlags.appId.Get() == nil {
		return ERROR_USAGE
	}

	call := service.App.Delete(appFlags.appId.String())
	app, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s", formatApp(app))

	out.Flush()
	return OK
}
