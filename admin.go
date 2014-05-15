package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/coreos/updatectl/client/update/v1"
	"github.com/coreos/updatectl/third_party/github.com/codegangsta/cli"
)

func AdminCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "admin-init",
			Usage:       "admin-init",
			Description: "Initializes the database.",
			Action:      adminInit,
		},
		{
			Name:        "admin-create-user",
			Usage:       "admin-create-user -u USER",
			Description: "Creates an admin user.",
			Action:      handle(adminCreateUser),
			Flags: []cli.Flag{
				cli.StringFlag{"user, u", "", "New Username"},
			},
		},
		{
			Name:        "admin-delete-user",
			Usage:       "admin-delete-user <username>",
			Description: "Deletes an admin user.",
			Action:      handle(adminDeleteUser),
		},
		{
			Name:        "admin-list-users",
			Usage:       "admin-list-users <username>",
			Description: "Lists admin user.",
			Action:      handle(adminListUsers),
		},
	}
}

func adminInit(c *cli.Context) {
	adminUrl := c.GlobalString("server") + "/admin/v1/init"
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
		os.Exit(1)
	}
}

func adminCreateUser(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	req := &update.AdminUserReq{
		UserName: c.String("user"),
	}
	call := service.Admin.CreateUser(req)
	u, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Token)
}

func adminDeleteUser(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	args := c.Args()
	if len(args) != 1 {
		cli.ShowCommandHelp(c, "admin-delete-user")
		os.Exit(1)
	}
	userName := args[0]
	call := service.Admin.DeleteUser(userName)
	u, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User %s deleted\n", u.User)
}

func adminListUsers(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	call := service.Admin.ListUsers()
	resp, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}

	for _, u := range resp.Users {
		fmt.Fprintf(out, "%s\n", u.User)
	}
	out.Flush()
}
