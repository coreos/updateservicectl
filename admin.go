package main

import (
	"fmt"
	"log"
	"text/tabwriter"

	"github.com/coreos/updatectl/client/update/v1"
)

var (
	cmdAdminUser = &Command{
		Name:    "admin-user",
		Usage:   "",
		Summary: "Operations for modifying admin users.",
		Subcommands: []*Command{
			cmdAdminUserCreate,
			cmdAdminUserList,
			cmdAdminUserDelete,
		},
	}
	cmdAdminUserCreate = &Command{
		Name:        "admin-user create",
		Usage:       "<username>",
		Description: "Creates an admin user.",
		Run:         adminUserCreate,
	}
	cmdAdminUserList = &Command{
		Name:        "admin-user list",
		Description: "Lists admin users.",
		Run:         adminUserList,
	}
	cmdAdminUserDelete = &Command{
		Name:        "admin-user delete",
		Usage:       "<username>",
		Description: "Deletes an admin user.",
		Run:         adminUserDelete,
	}
)

func adminUserCreate(args []string, service *update.Service, out *tabwriter.Writer) int {
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

func adminUserDelete(args []string, service *update.Service, out *tabwriter.Writer) int {
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

func adminUserList(args []string, service *update.Service, out *tabwriter.Writer) int {
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
