package downloader

type WebsiteDownloader struct{
	UrlList []string
	StorageLocation string 
}

func NewWebsiteDownloader(urlList []string, storageLocation string) *WebsiteDownloader{
	return &WebsiteDownloader{UrlList: urlList, StorageLocation: storageLocation} 
}


