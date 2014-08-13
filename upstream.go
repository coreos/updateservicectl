package main

import (
	"fmt"
	"log"
	"text/tabwriter"

	"github.com/coreos/updateservicectl/client/update/v1"
)

var (
	upstreamFlags struct {
		id    int64
		url   StringFlag
		label StringFlag
	}

	cmdUpstream = &Command{
		Name:        "upstream",
		Summary:     "Manage upstreams.",
		Description: `Synchronize published channels and packages from upstream installations of CoreOS Update Service.`,
		Subcommands: []*Command{
			cmdUpstreamCreate,
			cmdUpstreamUpdate,
			cmdUpstreamList,
			cmdUpstreamDelete,
			cmdUpstreamSync,
		},
	}
	cmdUpstreamCreate = &Command{
		Name:        "upstream create",
		Usage:       "[OPTION]...",
		Description: `Create a new upstream.`,
		Run:         upstreamCreate,
	}
	cmdUpstreamUpdate = &Command{
		Name:        "upstream update",
		Usage:       "[OPTION]...",
		Description: `Update an upstream.`,
		Run:         upstreamUpdate,
	}
	cmdUpstreamList = &Command{
		Name:        "upstream list",
		Description: `List all of the upstreams.`,
		Run:         upstreamList,
	}
	cmdUpstreamDelete = &Command{
		Name:        "upstream delete",
		Description: `Delete an upstream.`,
		Run:         upstreamDelete,
	}
	cmdUpstreamSync = &Command{
		Name:        "upstream sync",
		Description: `Sync all upstreams.`,
		Run:         upstreamSync,
	}
)

func init() {
	cmdUpstreamCreate.Flags.Var(&upstreamFlags.url, "url", "The root url of the upstream Update Service.")
	cmdUpstreamCreate.Flags.Var(&upstreamFlags.label, "label", "The label of the upstream Update Service.")

	cmdUpstreamUpdate.Flags.Int64Var(&upstreamFlags.id, "id", 0, "The id of the upstream to update.")
	cmdUpstreamUpdate.Flags.Var(&upstreamFlags.url, "url", "The root url of the upstream Update Service.")
	cmdUpstreamUpdate.Flags.Var(&upstreamFlags.label, "label", "The label of the upstream Update Service.")

	cmdUpstreamDelete.Flags.Int64Var(&upstreamFlags.id, "id", 0, "The id of the upstream to delete.")
}

func writeUpstreamHeading(out *tabwriter.Writer) {
	fmt.Fprintln(out, "Id\tUrl\tLabel")
}

func formatUpstream(us *update.Upstream) string {
	return fmt.Sprintf("%s\t%s\t%s\n", us.Id, us.Url, us.Label)
}

func upstreamCreate(args []string, service *update.Service, out *tabwriter.Writer) int {
	if upstreamFlags.url.Get() == nil {
		return ERROR_USAGE
	}

	req := &update.Upstream{
		Url:   upstreamFlags.url.String(),
		Label: upstreamFlags.label.String(),
	}
	call := service.Upstream.Insert(req)

	upstream, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}

	writeUpstreamHeading(out)
	fmt.Fprintf(out, "%s", formatUpstream(upstream))
	out.Flush()

	return OK
}

func upstreamUpdate(args []string, service *update.Service, out *tabwriter.Writer) int {
	if upstreamFlags.url.Get() == nil || upstreamFlags.id == 0 {
		return ERROR_USAGE
	}

	req := &update.Upstream{
		Id:    upstreamFlags.id,
		Url:   upstreamFlags.url.String(),
		Label: upstreamFlags.label.String(),
	}
	call := service.Upstream.Update(req.Id, req)

	upstream, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}

	writeUpstreamHeading(out)
	fmt.Fprintf(out, "%s", formatUpstream(upstream))
	out.Flush()

	return OK
}

func upstreamDelete(args []string, service *update.Service, out *tabwriter.Writer) int {
	if upstreamFlags.id == 0 {
		return ERROR_USAGE
	}

	call := service.Upstream.Delete(upstreamFlags.id)
	upstream, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}

	writeUpstreamHeading(out)
	fmt.Fprintf(out, "%s", formatUpstream(upstream))
	out.Flush()

	return OK
}

func upstreamList(args []string, service *update.Service, out *tabwriter.Writer) int {
	call := service.Upstream.List()
	upstreams, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}

	writeUpstreamHeading(out)
	for _, us := range upstreams.Items {
		fmt.Fprintf(out, "%s", formatUpstream(us))
	}
	out.Flush()

	return OK
}

func upstreamSync(args []string, service *update.Service, out *tabwriter.Writer) int {
	call := service.Upstream.Sync()
	resp, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "Status: %s\n", resp.Status)
	if resp.Detail != "" {
		fmt.Fprintf(out, "Detail: %s\n", resp.Detail)
	}
	out.Flush()

	return OK
}
