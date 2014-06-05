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
	cmdAdminInit = &Command{
		Name:        "admin-init",
		Description: "Initializes the database.",
		Run:         adminInit,
	}
	cmdAdminCreateUser = &Command{
		Name:        "admin-create-user",
		Usage:       "<username>",
		Description: "Creates an admin user.",
		Run:         adminCreateUser,
	}
	cmdAdminDeleteUser = &Command{
		Name:        "admin-delete-user",
		Usage:       "<username>",
		Description: "Deletes an admin user.",
		Run:         adminDeleteUser,
	}
	cmdAdminListUsers = &Command{
		Name:        "admin-list-users",
		Description: "Lists admin users.",
		Run:         adminListUsers,
	}
)

func adminInit(args []string, service *update.Service, out *tabwriter.Writer) int {
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

func adminCreateUser(args []string, service *update.Service, out *tabwriter.Writer) int {
	if len(args) != 1 {
		return ERROR_USAGE
	}

	req := &update.AdminUserReq{
		UserName: args[0],
	}
	call := service.Admin.CreateUser(req)
	u, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Token)
	return OK
}

func adminDeleteUser(args []string, service *update.Service, out *tabwriter.Writer) int {
	if len(args) != 1 {
		return ERROR_USAGE
	}
	userName := args[0]
	call := service.Admin.DeleteUser(userName)
	u, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User %s deleted\n", u.User)
	return OK
}

func adminListUsers(args []string, service *update.Service, out *tabwriter.Writer) int {
	call := service.Admin.ListUsers()
	resp, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}

	for _, u := range resp.Users {
		fmt.Fprintf(out, "%s\n", u.User)
	}
	out.Flush()
	return OK
}
