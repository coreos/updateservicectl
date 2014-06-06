package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/tabwriter"

	"github.com/coreos-inc/updatectl/client/update/v1"
)

var (
	cmdDatabase = &Command{
		Name:    "database",
		Usage:   "",
		Summary: "Operations for dealing with the backend database.",
		Subcommands: []*Command{
			cmdDatabaseInit,
		},
	}
	cmdDatabaseInit = &Command{
		Name:        "database init",
		Usage:       "",
		Description: "Initialize the database.",
		Run:         databaseInit,
	}
)

func databaseInit(args []string, service *update.Service, out *tabwriter.Writer) int {
	adminUrl := globalFlags.Server + "/admin/v1/init"
	client := &http.Client{}
	resp, err := client.Get(adminUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	if string(body) != "ok" {
		return ERROR_API
	}
	return OK
}
