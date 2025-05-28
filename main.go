package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	folder := flag.String("folder", "", "Path to the folder to zip")
	username := flag.String("username", "", "Username to authorize")
	password := flag.String("password", "", "Password to authorize")
	webdavUrl := flag.String("webdav-url", "", "WebDAV URL to upload to")

	flag.Parse()

	if *folder == "" {
		fmt.Println("Error: --folder is required")
		os.Exit(1)
	}

	if *username == "" {
		fmt.Println("Error: --username is required")
		os.Exit(1)
	}

	if *password == "" {
		fmt.Println("Error: --password is required")
		os.Exit(1)
	}

	if *webdavUrl == "" {
		fmt.Println("Error: --webdav-url is required")
		os.Exit(1)
	}

	archFilename := generateArchiveName()
	zipFilepath := filepath.Join(*folder, "..", archFilename)

	err := zipFolder(*folder, zipFilepath)
	if err != nil {
		fmt.Printf("Failed to zip folder: %v\n", err)
		os.Exit(1)
	}

	err = sendToWebdav(zipFilepath, fmt.Sprintf("%s/%s", *webdavUrl, archFilename), *username, *password)
	if err != nil {
		fmt.Printf("Failed to send to WebDAV: %v\n", err)
		os.Exit(1)
	}

	err = os.Remove(zipFilepath)
	if err != nil {
		fmt.Printf("Failed to remove zip file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Archive with name %s sent to %s\n", archFilename, *webdavUrl)
}
