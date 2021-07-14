package main

import (
	"flag"
	"strings"

	"github.com/mgould1799/WebsiteDownloader/downloader"
)

func main() {
	urlListString := flag.String("urls", "https://google.com,https://facebook.com", "a list of comma separated urls to download")
	fileStorageLocation := flag.String("storageLocation", "/temp", "path to store the downloaded sites")
	flag.Parse()

	// parse the url list
	// TODO: check if the string contains a , before splitting
	urlList := strings.Split(*urlListString, ",")

	websiteDownloader := downloader.NewWebsiteDownloader(urlList, *fileStorageLocation)
	websiteDownloader.Run()
}
