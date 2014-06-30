package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/coreos/updatectl/client/update/v1"
)

var (
	cmdDatabase = &Command{
		Name:    "database",
		Usage:   "",
		Summary: "Operations for dealing with the backend database.",
		Subcommands: []*Command{
			cmdDatabaseInit,
			cmdDatabaseBackup,
		},
	}
	cmdDatabaseInit = &Command{
		Name:        "database init",
		Usage:       "",
		Description: "Initialize the database.",
		Run:         databaseInit,
	}
	cmdDatabaseBackup = &Command{
		Name:        "database backup",
		Usage:       "<output file>",
		Description: "Grab a backup of the database.",
		Run:         databaseBackup,
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

func databaseBackup(args []string, service *update.Service, out *tabwriter.Writer) int {
	if len(args) != 1 {
		return ERROR_USAGE
	}
	backupUrl := globalFlags.Server + "/db/backup"
	client := getHawkClient(globalFlags.User, globalFlags.Key)
	resp, err := client.Get(backupUrl)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Fatal(string(body))
	}
	defer resp.Body.Close()
	outFile, err := os.Create(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return OK
}
