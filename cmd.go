package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/coreos-inc/updatectl/auth"
	"github.com/coreos-inc/updatectl/client/update/v1"
	"github.com/coreos-inc/updatectl/version"
)

const (
	cliName        = "updatectl"
	cliDescription = "updatectl is a command line driven interface to the roller."

	OK = iota
	// Error Codes
	ERROR_API
	ERROR_USAGE
	ERROR_NO_COMMAND
)

type Command struct {
	Name        string       // Name of the Command and the string to use to invoke it
	Summary     string       // One-sentence summary of what the Command does
	Usage       string       // Usage options/arguments
	Description string       // Detailed description of command
	Flags       flag.FlagSet // Set of flags associated with this command
	Run         handlerFunc  // Run a command with the given arguments
}

var (
	out           *tabwriter.Writer
	globalFlagSet *flag.FlagSet
	commands      []*Command

	globalFlags struct {
		Server  string
		User    string
		Key     string
		Debug   bool
		Version bool
	}
)

func init() {
	out = new(tabwriter.Writer)
	out.Init(os.Stdout, 0, 8, 1, '\t', 0)

	globalFlagSet = flag.NewFlagSet(cliName, flag.ExitOnError)
	globalFlagSet.StringVar(&globalFlags.Server, "server", "http://localhost:8000", "Update server to connect to")
	globalFlagSet.BoolVar(&globalFlags.Debug, "debug", false, "Output debugging info to stderr")
	globalFlagSet.BoolVar(&globalFlags.Version, "version", false, "Print version information")
	globalFlagSet.StringVar(&globalFlags.User, "user", os.Getenv("UPDATECTL_USER"), "API Username")
	globalFlagSet.StringVar(&globalFlags.Key, "key", os.Getenv("UPDATECTL_KEY"), "API Key")

	commands = []*Command{
		// admin.go
		cmdAdminInit,
		cmdAdminCreateUser,
		cmdAdminDeleteUser,
		cmdAdminListUsers,
		// app.go
		cmdListApps,
		cmdCreateApp,
		cmdUpdateApp,
		cmdDeleteApp,
		// channel.go
		cmdListChannels,
		cmdUpdateChannel,
		// client_update.go
		cmdListClientUpdates,
		cmdListAppVersions,
		// fakeclients.go
		cmdFakeClients,
		// group.go
		cmdListGroups,
		cmdNewGroup,
		cmdDeleteGroup,
		cmdUpdateGroup,
		cmdPauseGroup,
		cmdUnpauseGroup,
		cmdRollupGroupVersions,
		cmdRollupGroupEvents,
		// help.go
		cmdHelp,
		// pkg.go
		cmdListPackages,
		cmdNewPackage,
		// watch.go
		cmdWatch,
	}
}

type handlerFunc func([]string, *update.Service, *tabwriter.Writer) int

func getHawkClient(user string, key string) *http.Client {
	return &http.Client{Transport: &auth.HawkRoundTripper{user, key}}
}

func handle(fn handlerFunc) func(f *flag.FlagSet) int {
	return func(f *flag.FlagSet) (exit int) {
		user := globalFlags.User
		key := globalFlags.Key
		client := getHawkClient(user, key)

		service, err := update.New(client)
		if err != nil {
			log.Fatal(err)
		}

		service.BasePath = globalFlags.Server + "/_ah/api/update/v1/"
		exit = fn(f.Args(), service, out)
		return
	}
}

func printVersion(out *tabwriter.Writer) {
	fmt.Fprintf(out, "%s version %s\n", cliName, version.Version)
	out.Flush()
}

func getAllFlags() (flags []*flag.Flag) {
	return getFlags(globalFlagSet)
}

func getFlags(flagset *flag.FlagSet) (flags []*flag.Flag) {
	flags = make([]*flag.Flag, 0)
	flagset.VisitAll(func(f *flag.Flag) {
		flags = append(flags, f)
	})
	return
}

func main() {
	globalFlagSet.Parse(os.Args[1:])
	var args = globalFlagSet.Args()

	if globalFlags.Version {
		printVersion(out)
		os.Exit(OK)
	}

	// no command specified - trigger help
	if len(args) < 1 {
		args = append(args, "help")
	}

	var cmd *Command

	// determine which Command should be run
	for _, c := range commands {
		if c.Name == args[0] {
			cmd = c
			if err := c.Flags.Parse(args[1:]); err != nil {
				fmt.Println(err.Error())
				os.Exit(ERROR_USAGE)
			}
			break
		}
	}

	if cmd == nil {
		fmt.Printf("%v: unknown subcommand: %q\n", cliName, args[0])
		fmt.Printf("Run '%v help' for usage.\n", cliName)
		os.Exit(ERROR_NO_COMMAND)
	}

	exit := handle(cmd.Run)(&cmd.Flags)

	if exit == ERROR_USAGE {
		printCommandUsage(cmd)
	}
	os.Exit(exit)
}
