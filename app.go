package main

import (
	"fmt"
	"log"
	"text/tabwriter"

	"github.com/coreos/updateservicectl/Godeps/_workspace/src/code.google.com/p/go-uuid/uuid"
	"github.com/coreos/updateservicectl/client/update/v1"
)

var (
	appFlags struct {
		appId       StringFlag
		label       StringFlag
		description StringFlag
	}

	cmdApp = &Command{
		Name:    "app",
		Summary: "Manage applications.",
		Subcommands: []*Command{
			cmdAppCreate,
			cmdAppList,
			cmdAppUpdate,
			cmdAppDelete,
		},
	}
	cmdAppCreate = &Command{
		Name:        "app create",
		Usage:       "[OPTION]...",
		Description: `Create a new application.`,
		Run:         appCreate,
	}
	cmdAppList = &Command{
		Name:        "app list",
		Description: `List all of the apps that exist including their label, token and update state.`,
		Run:         appList,
	}
	cmdAppUpdate = &Command{
		Name:        "app update",
		Usage:       "[OPTION]...",
		Description: `Update an app's label or description.`,
		Run:         appUpdate,
	}
	cmdAppDelete = &Command{
		Name:        "app delete",
		Usage:       "[OPTION]...",
		Description: `Delete an app.`,
		Run:         appDelete,
	}
)

func init() {
	cmdAppCreate.Flags.Var(&appFlags.appId, "app-id", "Application UUID. If not provided, one will be randomly generated.")
	cmdAppCreate.Flags.Var(&appFlags.label, "label", "New application label.")
	cmdAppCreate.Flags.Var(&appFlags.description, "description", "New application description.")

	cmdAppUpdate.Flags.Var(&appFlags.appId, "app-id", "Application ID to update.")
	cmdAppUpdate.Flags.Var(&appFlags.label, "label", "Set application label.")
	cmdAppUpdate.Flags.Var(&appFlags.description, "description", "Set application description.")

	cmdAppDelete.Flags.Var(&appFlags.appId, "app-id", "Application ID to delete.")
}

const appHeader = "Id\tLabel\tDescription\n"

func formatApp(app *update.App) string {
	return fmt.Sprintf("%s\t%s\t%s\n", app.Id, app.Label, app.Description)
}

func appList(args []string, service *update.Service, out *tabwriter.Writer) int {
	listCall := service.App.List()
	list, err := listCall.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(out, appHeader)
	for _, app := range list.Items {
		fmt.Fprintf(out, "%s", formatApp(app))
	}

	out.Flush()
	return OK
}

func appCreate(args []string, service *update.Service, out *tabwriter.Writer) int {
	if appFlags.appId.Get() == nil {
		appFlags.appId.Set(uuid.New())
	}

	appReq := &update.AppInsertReq{
		Id:          appFlags.appId.String(),
		Label:       appFlags.label.String(),
		Description: appFlags.description.String(),
	}
	call := service.App.Insert(appReq)
	app, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(out, appHeader)
	fmt.Fprintf(out, "%s", formatApp(app))

	out.Flush()
	return OK

}

func appUpdate(args []string, service *update.Service, out *tabwriter.Writer) int {
	if appFlags.appId.Get() == nil || appFlags.label.Get() == nil || appFlags.description.Get() == nil {
		return ERROR_USAGE
	}

	appReq := &update.AppUpdateReq{Label: appFlags.label.String(), Description: appFlags.description.String()}
	call := service.App.Update(appFlags.appId.String(), appReq)
	app, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(out, appHeader)
	fmt.Fprintf(out, "%s", formatApp(app))

	out.Flush()
	return OK

}

func appDelete(args []string, service *update.Service, out *tabwriter.Writer) int {
	if appFlags.appId.Get() == nil {
		return ERROR_USAGE
	}

	call := service.App.Delete(appFlags.appId.String())
	app, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(out, appHeader)
	fmt.Fprintf(out, "%s", formatApp(app))

	out.Flush()
	return OK
}
