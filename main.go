package main

import (
	"flag"
	"strings"
)

func main() {
	urlListString := flag.String("urls", "google.com,facebook.com", "a list of comma separated urls to download")
	fileStorageLocation := flag.String("storageLocation", "/temp", "path to store the downloaded sites")
	flag.Parse()

	urlList := strings.Split(*urlListString, ",")

	_ = urlList
	_ = fileStorageLocation

	// websiteDownloader := downloader.NewWebsiteDownloader(urlList, fileStorageLocation)
	// websiteDownloader.Run()

}
