package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mgould1799/WebsiteDownloader/downloader"
)

func main() {
	fmt.Println("hi")
	urlListString := flag.String("urls", "google.com,facebook.com", "a list of comma separated urls to download")
	fileStorageLocation := flag.String("storageLocation", "/temp", "path to store the downloaded sites")
	flag.Parse()

	// parse the url list
	// TODO: check if the string contains a , before splitting 
	urlList := strings.Split(*urlListString, ",")

	websiteDownloader := downloader.NewWebsiteDownloader(urlList, *fileStorageLocation)
	websiteDownloader.Run()
}
