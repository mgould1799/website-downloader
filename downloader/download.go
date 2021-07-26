package downloader

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// a struct to host the storage location and urllist
type WebsiteDownloader struct {
	UrlList         []string
	StorageLocation string
	MaxDownloads    int
}

// a funciton to create a new object on the string
func NewWebsiteDownloader(urlList []string, storageLocation string) *WebsiteDownloader {
	return &WebsiteDownloader{UrlList: urlList, StorageLocation: storageLocation, MaxDownloads: 2}
}

// run the actual downloading of the urls
func (wd *WebsiteDownloader) Run() {
	// check if storage location is empty
	err := wd.checkDirectory()
	if err != nil {
		panic(err)
	}
	// put the urls into a channel
	urlChannel := wd.outputUrls()
	// grab those  vars from a channel and download them concurrently
	creationChannel := wd.downloadUrls(urlChannel)

	for creation := range creationChannel {
		if creation.Successful != "" {
			fmt.Println(creation.Successful)
		}
		if creation.ErrorMsg != "" {
			fmt.Println(creation.ErrorMsg)
		}
	}
}

// puts the urls into a channel
func (wd *WebsiteDownloader) outputUrls() <-chan string {
	urlChannel := make(chan string, len(wd.UrlList))
	defer close(urlChannel)
	for _, url := range wd.UrlList {
		// add https:// to a url for incase it needs it
		if !strings.Contains(url, "https://") {
			temp := url
			url = fmt.Sprintf("https://%s", temp)
		}
		urlChannel <- url
	}
	return urlChannel
}

// downlaods the sites and saves them to a place
// output response to channel
// constraint on the number of go routines that can be ran
func (wd *WebsiteDownloader) downloadUrls(urls <-chan string) <-chan Creation {
	// create the output channel
	output := make(chan Creation)

	// create a wait group to control the number of workers
	var wg sync.WaitGroup
	// loop through the number of workers allowed
	for i := 0; i < wd.MaxDownloads; i++ {
		// say hey wait
		wg.Add(1)
		// span off on the go routine
		go func() {
			defer wg.Done()
			for url := range urls {
				// download the site
				downloadedSite, err := downloadSite(url)
				if err != nil {
					output <- *newCreation("", fmt.Sprintf("%v - error occurred while downloading site - %v", url, err))
					break
				}

				// create the location to save the string
				locationString, err := wd.createLocationString(url)
				if err != nil {
					output <- *newCreation("", fmt.Sprintf("%v - error occurred while creating location string - %v", url, err))
					break
				}
				// save the site to specified location
				err = saveSite(downloadedSite, locationString)
				if err != nil {
					output <- *newCreation("", fmt.Sprintf("%v - error occurred while saving site - %v", url, err))
					break
				}
				// output to the channel
				output <- *newCreation(fmt.Sprintf("successfully downloaded - %v - saved to %v", url, locationString), "")
			}
		}()

	}

	// clean up the output channel to return it
	go func() {
		wg.Wait()
		close(output)
	}()
	return output
}

// check if the give directory exists if not create it
func (wd *WebsiteDownloader) checkDirectory() error {
	newpath := filepath.Join(".", wd.StorageLocation)
	return os.MkdirAll(newpath, os.ModePerm)
}

// create name of the file
func (wd *WebsiteDownloader) createLocationString(urlString string) (string, error) {
	fileUrl, err := url.ParseRequestURI(urlString)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s.html", wd.StorageLocation, fileUrl.Host), nil
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

// downloads a site
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
