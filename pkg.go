package main

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/coreos/updatectl/client/update/v1"
	"github.com/coreos/updatectl/third_party/github.com/codegangsta/cli"
)

type MetadataFile struct {
	MetadataSize         string `json:"metadata_size"`
	MetadataSignatureRsa string `json:"metadata_signature_rsa"`
}

func PackageCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "list-packages",
			Usage:       "list-packages <appId>",
			Description: `List all of the packages that exist including their metadata.`,
			Action:      handle(listPackages),
		},
		{
			Name:        "new-package",
			Usage:       "new-package [OPTION]... <appId>",
			Description: `Create a new package given version,  meta file.`,
			Action:      handle(newPackage),
			Flags: []cli.Flag{
				cli.StringFlag{"version", "", ""},
				cli.StringFlag{"url", "", ""},
				cli.StringFlag{"file", "update.gz", ""},
				cli.StringFlag{"meta", "update.meta", ""},
			},
		},
	}
}

func newPackage(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	args := c.Args()

	if len(args) != 1 {
		fmt.Println("usage: <appid>")
		os.Exit(1)
	}

	file := c.String("meta")
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("reading %s failed: %v", file, err)
	}
	var meta MetadataFile
	err = json.Unmarshal(content, &meta)
	if err != nil {
		log.Fatalf("reading %s failed: %v", file, err)
	}

	file = c.String("file")
	info, err := os.Stat(file)
	if err != nil {
		log.Fatalf("state of %s failed: %v", file, err)
	}
	content, err = ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("reading %s failed: %v", file, err)
	}

	var sha1base64 bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &sha1base64)
	sha1h := sha1.New()
	io.Copy(sha1h, bytes.NewReader(content))
	encoder.Write(sha1h.Sum(nil))
	encoder.Close()

	var sha256base64 bytes.Buffer
	encoder = base64.NewEncoder(base64.StdEncoding, &sha256base64)
	sha256h := sha256.New()
	io.Copy(sha256h, bytes.NewReader(content))
	encoder.Write(sha256h.Sum(nil))
	encoder.Close()

	pkg := &update.Package{
		Url:                  c.String("url"),
		Size:                 strconv.FormatInt(info.Size(), 10),
		Sha1Sum:              sha1base64.String(),
		Sha256Sum:            sha256base64.String(),
		MetadataSignatureRsa: meta.MetadataSignatureRsa,
		MetadataSize:         meta.MetadataSize,
	}

	jbytes, _ := json.MarshalIndent(pkg, "", " ")
	fmt.Printf("%s\n", string(jbytes))

	call := service.App.Package.Insert(args[0], c.String("version"), pkg)
	pkg, err = call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(out, args[0], c.String("version"))

	out.Flush()
}

func listPackages(c *cli.Context, service *update.Service, out *tabwriter.Writer) {
	args := c.Args()

	if len(args) != 1 {
		fmt.Println("usage: <appid>")
		os.Exit(1)
	}

	call := service.App.Package.List(args[0])
	list, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(out, "Version\t\tURL\tSize")
	for _, pkg := range list.Items {
		fmt.Fprintf(out, "%s\t%s\t%s\n", pkg.Version, pkg.Url, pkg.Size)
		out.Flush()
	}
}
