package downloader

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// a struct to host the storage location and urllist
type WebsiteDownloader struct {
	UrlList         []string
	StorageLocation string
}

// a funciton to create a new object on the string
func NewWebsiteDownloader(urlList []string, storageLocation string) *WebsiteDownloader {
	return &WebsiteDownloader{UrlList: urlList, StorageLocation: storageLocation}
}

// run the actual downloading of the urls
func (wd *WebsiteDownloader) Run() {
	// put the urls into a channel
	urlChannel := wd.outputUrls()
	// grab those  vars from a channel and download them concurrently
	wd.downloadUrls(urlChannel)
}

// downloads a site a site
// returns a bytes.buffer and error 
func downloadSite(url string) (*bytes.Buffer, error) {
	// run a get request for the website
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// put the response in a buffer
	defer resp.Body.Close()
	buffer := bytes.NewBuffer(nil)
	io.Copy(buffer, resp.Body)

	// return the buffer and nil since it was successful 
	return buffer, nil
}

// puts the urls into a channel
func (wd *WebsiteDownloader) outputUrls() <- chan string{
	urlChannel := make(chan string, len(wd.UrlList))
	defer close(urlChannel)
	for _, url := range wd.UrlList{
		urlChannel <- url 
	}
	return urlChannel
}

// downlaods the sites and aves them to a place
func (wd *WebsiteDownloader) downloadUrls(urls <- chan string) error{
	for url := range urls {
		// download the site
		downloadedSite, err := downloadSite(url)
		if err != nil {
			return err
		}

		locationString, err := createLocationString(url)
		if err != nil {
			return nil 
		}
		err = saveSite(downloadedSite, locationString)
		if err != nil {
			return nil 
		}
	}

	return nil 
}

// create name of the file 
func createLocationString(urlString string) (string, error) {
	fileUrl, err := url.ParseRequestURI(urlString)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("./temp/%s.html", fileUrl.Host), nil 
}

// saves the site to a specified location 
// take in a bytes bugger and a string to save the site too 
// returns an error if any occur
func saveSite(downloadedSite *bytes.Buffer, saveLocation string) error {
	// create the file. The directory needs to be created first
	file, err := os.Create(saveLocation)
	if err != nil {
		return err
	}
	defer file.Close()

	// copy the bytes buffer into the downloaded site
	_, err = io.Copy(file, downloadedSite)
	if err != nil {
		return err
	}

	// return nil since no error occurred
	return nil
}
