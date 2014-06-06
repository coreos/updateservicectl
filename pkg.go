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

	"github.com/coreos-inc/updatectl/client/update/v1"
)

type MetadataFile struct {
	MetadataSize         string `json:"metadata_size"`
	MetadataSignatureRsa string `json:"metadata_signature_rsa"`
}

var (
	packageFlags struct {
		appId   StringFlag
		version StringFlag
		url     string
		file    string
		meta    string
	}

	cmdPackage = &Command{
		Name:    "package",
		Summary: "List or create packages for an application.",
		Subcommands: []*Command{
			cmdPackageList,
			cmdPackageCreate,
		},
	}

	cmdPackageList = &Command{
		Name:        "package list",
		Usage:       "[OPTION]...",
		Description: `List all of the packages that exist including their metadata.`,
		Run:         packageList,
	}
	cmdPackageCreate = &Command{
		Name:        "package create",
		Usage:       "[OPTION]...",
		Description: `Create a new package for an application.`,
		Run:         packageCreate,
	}
)

func init() {
	cmdPackageList.Flags.Var(&packageFlags.appId, "app-id", "Application to list the package of.")

	cmdPackageCreate.Flags.Var(&packageFlags.appId, "app-id", "Application to add the package to.")
	cmdPackageCreate.Flags.Var(&packageFlags.version, "version", "Application version associated with the package.")
	cmdPackageCreate.Flags.StringVar(&packageFlags.url, "url", "", "Package URL.")
	cmdPackageCreate.Flags.StringVar(&packageFlags.file, "file", "update.gz", "Package file.")
	cmdPackageCreate.Flags.StringVar(&packageFlags.meta, "meta", "", "JSON file containing metadata.")
}

func packageCreate(args []string, service *update.Service, out *tabwriter.Writer) int {
	if packageFlags.appId.Get() == nil ||
		packageFlags.version.Get() == nil {
		return ERROR_USAGE
	}

	file := packageFlags.file
	info, err := os.Stat(file)
	if err != nil {
		log.Fatalf("stat of %s failed: %v", file, err)
	}
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("reading %s failed: %v", file, err)
	}

	file = packageFlags.meta
	var meta MetadataFile
	if file != "" {
		content, err = ioutil.ReadFile(file)
		if err != nil {
			log.Fatalf("reading %s failed: %v", file, err)
		}
		err = json.Unmarshal(content, &meta)
		if err != nil {
			log.Fatalf("reading %s failed: %v", file, err)
		}
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
		Url:                  packageFlags.url,
		Size:                 strconv.FormatInt(info.Size(), 10),
		Sha1Sum:              sha1base64.String(),
		Sha256Sum:            sha256base64.String(),
		MetadataSignatureRsa: meta.MetadataSignatureRsa,
		MetadataSize:         meta.MetadataSize,
	}

	jbytes, _ := json.MarshalIndent(pkg, "", " ")
	fmt.Printf("%s\n", string(jbytes))

	call := service.App.Package.Insert(packageFlags.appId.String(), packageFlags.version.String(), pkg)
	pkg, err = call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(out, packageFlags.appId, packageFlags.version)

	out.Flush()
	return OK
}

func packageList(args []string, service *update.Service, out *tabwriter.Writer) int {
	if packageFlags.appId.Get() == nil {
		return ERROR_USAGE
	}

	call := service.App.Package.List(packageFlags.appId.String())
	list, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(out, "Version\tURL\tSize")
	for _, pkg := range list.Items {
		fmt.Fprintf(out, "%s\t%s\t%s\n", pkg.Version, pkg.Url, pkg.Size)
	}

	out.Flush()
	return OK
}
