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
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"text/tabwriter"

	"github.com/coreos/updateservicectl/Godeps/_workspace/src/github.com/cheggaaa/pb"
	update "github.com/coreos/updateservicectl/client/update/v1"
)

type MetadataFile struct {
	MetadataSize         string `json:"metadata_size"`
	MetadataSignatureRsa string `json:"metadata_signature_rsa"`
}

var (
	downloadGroup   sync.WaitGroup
	createBulkGroup sync.WaitGroup

	packageFlags struct {
		appId        StringFlag
		version      StringFlag
		url          string
		file         string
		meta         string
		releaseNotes string
		saveDir      string
		bulkDir      string
		baseUrl      string
	}

	cmdPackage = &Command{
		Name:    "package",
		Summary: "List or create packages for an application.",
		Subcommands: []*Command{
			cmdPackageList,
			cmdPackageCreate,
			cmdPackageDelete,
			cmdPackageDownload,
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
		Subcommands: []*Command{
			cmdPackageCreateBulk,
		},
	}
	cmdPackageDelete = &Command{
		Name:        "package delete",
		Usage:       "[OPTION]...",
		Description: `Delete a package for an application.`,
		Run:         packageDelete,
	}
	cmdPackageDownload = &Command{
		Name:        "package download",
		Usage:       "[OPTION]...",
		Description: `Download published packages to local disk.`,
		Run:         packageDownload,
	}
	cmdPackageCreateBulk = &Command{
		Name:        "package create bulk",
		Usage:       "[OPTION]...",
		Description: `Upload package from a folder output by 'package donload'.`,
		Run:         packageCreateBulk,
	}
)

func init() {
	cmdPackageList.Flags.Var(&packageFlags.appId, "app-id",
		"Application to list the package of.")

	cmdPackageCreate.Flags.Var(&packageFlags.appId, "app-id",
		"Application to add the package to.")
	cmdPackageCreate.Flags.Var(&packageFlags.version, "version",
		"Application version associated with the package.")
	cmdPackageCreate.Flags.StringVar(&packageFlags.url, "url", "",
		"Package URL.")
	cmdPackageCreate.Flags.StringVar(&packageFlags.file, "file",
		"update.gz", "Package file.")
	cmdPackageCreate.Flags.StringVar(&packageFlags.meta, "meta", "",
		"JSON file containing metadata.")
	cmdPackageCreate.Flags.StringVar(&packageFlags.releaseNotes,
		"release-notes", "",
		"File contianing release notes for package.")

	cmdPackageCreateBulk.Flags.StringVar(&packageFlags.bulkDir,
		"dir", "",
		"Directory containing files to upload.")
	cmdPackageCreateBulk.Flags.StringVar(&packageFlags.baseUrl,
		"base-url", "",
		"URL base packages are stored at.")

	cmdPackageDelete.Flags.Var(&packageFlags.appId, "app-id",
		"Application with package to delete.")
	cmdPackageDelete.Flags.Var(&packageFlags.version, "version",
		"Version of package to delete.")

	cmdPackageDownload.Flags.StringVar(&packageFlags.saveDir, "dir",
		"", "Directory to save downloaded packages in.")
}

func formatPackage(pkg *update.Package) string {
	return fmt.Sprintf("%s\t%s\t%s\n", pkg.Version, pkg.Url, pkg.Size)
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

	metaFile := packageFlags.meta
	var meta MetadataFile
	if metaFile != "" {
		content, err = ioutil.ReadFile(metaFile)
		if err != nil {
			log.Fatalf("reading %s failed: %v", metaFile, err)
		}
		err = json.Unmarshal(content, &meta)
		if err != nil {
			log.Fatalf("reading %s failed: %v", metaFile, err)
		}
	}

	releaseNotesFile := packageFlags.releaseNotes
	var notes = make([]byte, 0)
	if releaseNotesFile != "" {
		notes, err = ioutil.ReadFile(releaseNotesFile)
		if err != nil {
			log.Fatalf("reading %s failed: %v", releaseNotesFile, err)
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
		ReleaseNotes:         string(notes),
	}

	jbytes, _ := json.MarshalIndent(pkg, "", " ")
	fmt.Printf("%s\n", string(jbytes))

	call := service.App.Package.Insert(packageFlags.appId.String(), packageFlags.version.String(), pkg)
	pkg, err = call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(out, packageFlags.appId.String(), packageFlags.version.String())

	out.Flush()
	return OK
}

func packageCreateBulk(args []string, service *update.Service, out *tabwriter.Writer) int {
	bulkDir := packageFlags.bulkDir
	if bulkDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			log.Print(err)
			return ERROR_USAGE
		}
		bulkDir = cwd
	}

	files, err := ioutil.ReadDir(bulkDir)
	if err != nil {
		log.Print(err)
		return ERROR_USAGE
	}

	var total int
	var errors int
	errorHandler := func(err error) {
		if err != nil {
			errors++
			log.Printf("Error while creating package. Error=%s", err)
			createBulkGroup.Done()
		}
	}

	for _, file := range files {
		if file.Mode().IsRegular() && strings.HasSuffix(file.Name(), "info.json") {
			total++
			createBulkGroup.Add(1)
			go createPackageFromInfoFile(
				path.Join(bulkDir, file.Name()),
				service,
				errorHandler,
			)
		}
	}
	createBulkGroup.Wait()
	log.Printf("Package metadata uploaded. Total=%d Errors=%d", total, errors)
	if packageFlags.baseUrl != "" {
		log.Printf("Please upload payloads to %s.", packageFlags.baseUrl)
	}
	return OK
}

func createPackageFromInfoFile(filename string, service *update.Service, handleError func(error)) {
	// Load metadata from package info.json into struct
	pkg := new(update.Package)
	jsonBody, err := ioutil.ReadFile(filename)
	if err != nil {
		handleError(err)
		return
	}

	err = json.Unmarshal(jsonBody, &pkg)
	if err != nil {
		handleError(err)
		return
	}

	log.Printf("Creating package with AppId=%s and Version=%s", pkg.AppId, pkg.Version)

	// If --base-url specified, rewrite hosting URL
	baseUrl := packageFlags.baseUrl
	if baseUrl != "" {
		filename := fmt.Sprintf(
			"%s_%s_%s",
			pkg.AppId, pkg.Version,
			path.Base(pkg.Url),
		)
		pkg.Url = path.Join(baseUrl, filename)
	}

	// Add package
	call := service.App.Package.Insert(pkg.AppId, pkg.Version, pkg)
	pkg, err = call.Do()

	if err != nil {
		handleError(err)
		return
	}
	createBulkGroup.Done()
	return
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

func packageDelete(args []string, service *update.Service, out *tabwriter.Writer) int {
	if packageFlags.appId.Get() == nil ||
		packageFlags.version.Get() == nil {
		return ERROR_USAGE
	}

	call := service.App.Package.Delete(packageFlags.appId.String(), packageFlags.version.String())
	pkg, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "%s", formatPackage(pkg))

	out.Flush()
	return OK

}

func packageDownload(args []string, service *update.Service, out *tabwriter.Writer) int {
	saveDir, err := getPackageSaveDirectory()
	if err != nil {
		log.Print(err)
		return ERROR_USAGE
	}

	call := service.App.Package.PublicList()
	pkgs, err := call.Do()
	if err != nil {
		log.Print(err)
		return ERROR_USAGE
	}

	// Setup progress bar
	var totalSize int64
	var totalPackages int
	for _, item := range pkgs.Items {
		for _, pkg := range item.Packages {
			pkgSize, err := strconv.ParseInt(pkg.Size, 10, 64)
			if err != nil {
				log.Println("Parse of package size failed.")
			}
			totalSize += pkgSize
			totalPackages++
		}
	}

	bar := pb.New64(totalSize).SetUnits(pb.U_BYTES)

	// Error handle for download worker
	errorHandler := func(pkg *update.Package) func(error) {
		return func(err error) {
			if err != nil {
				fmt.Fprintf(os.Stderr,
					"Error while downloading. AppId=%s, Version=%s, URL=%s, Error=%s",
					pkg.AppId, pkg.Version, pkg.Url, err,
				)

				if err = os.Remove(path.Join(saveDir, path.Base(pkg.Url))); err != nil {
					fmt.Fprintf(os.Stderr,
						"Tried to remove file and got error: %s",
						err,
					)
				}

				downloadGroup.Done()
			}
		}
	}

	// Download package payloads in parallel
	log.Printf("Downloading %d packages.", totalPackages)
	bar.Start()
	for _, item := range pkgs.Items {
		for _, pkg := range item.Packages {
			downloadGroup.Add(1)
			go downloadPackagePayload(pkg, saveDir, bar, errorHandler(pkg))
		}
	}
	downloadGroup.Wait()
	bar.Finish()
	return OK
}

func downloadPackagePayload(pkg *update.Package, saveTo string, bar *pb.ProgressBar, handle func(error)) {
	// Ensure we have a valid package URL
	pkgUrl, err := url.Parse(pkg.Url)
	if err != nil {
		handle(err)
		return
	}

	// Currently only supports files hosted publicly on HTTP/HTTPS
	if pkgUrl.Scheme != "http" && pkgUrl.Scheme != "https" {
		err = fmt.Errorf("Cannot download package with scheme %s", pkgUrl.Scheme)
		handle(err)
		return
	}

	// Download the package
	res, err := http.Get(pkg.Url)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		handle(err)
		return
	}

	// Save the file by Applciation, Version, and Filename
	filename := fmt.Sprintf("%s_%s_%s", pkg.AppId, pkg.Version, path.Base(pkg.Url))
	out, err := os.Create(path.Join(saveTo, filename))
	if out != nil {
		defer out.Close()
	}

	if err != nil {
		handle(err)
		return
	}

	// We will hash the file as we download it.
	sha1h := sha1.New()
	// Write to file, hash, and progress bar.
	n, err := io.Copy(io.MultiWriter(out, sha1h, bar), res.Body)
	if err != nil {
		handle(err)
		return
	}

	// Verify downloaded size matches the package's size.
	pkgSize, err := strconv.ParseInt(pkg.Size, 10, 64)
	if n != pkgSize {
		err = fmt.Errorf("Download size does not match package size. %d != %d", n, pkgSize)
		handle(err)
		return
	}

	// Decode the SHA1 sum in the package manifest.
	pkgSha1, err := base64.StdEncoding.DecodeString(pkg.Sha1Sum)
	if err != nil {
		handle(err)
		return
	}
	// Verify the hash matches
	if string(pkgSha1) != string(sha1h.Sum(nil)) {
		err = fmt.Errorf("SHA1 sums do not match: %s != $s", string(pkgSha1), string(sha1h.Sum(nil)))
		handle(err)
		return
	}
	// Write out an info.json file containing metadata
	err = writePackageInfo(pkg, saveTo)
	if err != nil {
		handle(err)
		return
	}

	downloadGroup.Done()
}

func getPackageSaveDirectory() (string, error) {
	saveDir := packageFlags.saveDir
	if saveDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		saveDir = cwd
	}

	info, err := os.Stat(saveDir)
	if err == nil {
		if info.IsDir() == false {
			return "", fmt.Errorf("%s is not a directory", saveDir)
		}
	} else if os.IsNotExist(err) {
		err := os.Mkdir(saveDir, 0700)
		if err != nil {
			return "", err
		}
	}
	return saveDir, nil
}

func writePackageInfo(pkg *update.Package, saveTo string) error {
	filename := fmt.Sprintf("%s_%s_info.json", pkg.AppId, pkg.Version)
	out, err := os.Create(path.Join(saveTo, filename))
	if out != nil {
		defer out.Close()
	}
	if err != nil {
		return err
	}

	output, err := json.Marshal(pkg)
	if err != nil {
		return err
	}

	out.Write(output)
	return nil
}
