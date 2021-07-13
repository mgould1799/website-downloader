package downloader

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

type WebsiteDownloader struct{
	UrlList []string
	StorageLocation string 
}

func NewWebsiteDownloader(urlList []string, storageLocation string) *WebsiteDownloader{
	return &WebsiteDownloader{UrlList: urlList, StorageLocation: storageLocation} 
}

func (wd *WebsiteDownloader) Run() {
	// constraint: this needs to be concurrent
	// create a method to download 
	// create a method to store the site in a location

	output, err := downloadSite("https://google.com")
	if err != nil {
		panic(err)
	}

	err = saveSite(output, "./temp/googledotcom.html")
	if err != nil {
		fmt.Println(err)
	}
}

func downloadSite(url string) (*bytes.Buffer, error){
	 resp, err := http.Get(url)
	 if err != nil {
		return nil, err 
	 }
	 // put the response in a buffer
	 defer resp.Body.Close()
	 buffer := bytes.NewBuffer(nil)
	 io.Copy(buffer, resp.Body)

	 return buffer, nil 
}

func saveSite(downloadedSite *bytes.Buffer, saveLocation string) error{
	file, err := os.Create(saveLocation)
	if err != nil {
		return err 
	}
	defer file.Close()

	
	_, err = io.Copy(file, downloadedSite)
	if err != nil {
		return nil 
	}

	return nil 
}