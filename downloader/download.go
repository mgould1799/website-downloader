package downloader

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func (wd *WebsiteDownloader) Run() {
	// constraint: this needs to be concurrent
	// create a method to download 
	// create a method to store the site in a location

	output, err := downloadSite("https://google.com")
	if err != nil {
		panic(err)
	}

	saveSite(output, "./temp/google.com")
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
	return nil 
}