# Website Downloader

## Description

This program is a simple program to download sites and save them to a directoy. This is to showcase a simple implementation of the pipeline concurrency pattern in golang.

## Getting Started

### Dependencies

* golang needs to be installed 
* all package dependencies can be found within the gomod file
* a directory created called temp


### Executing program

```
mkdir temp 
go mod download
go build 
./WebsiteDownloader
```

####Flags that can be passed

`--urls`
Urls is a list of comma separated urls to pull down and save

example run 
```
./WebsiteDownloader --urls=https://google.com,https://facebook.com
````

## Acknowledgments

Inspiration, code snippets, etc.
* https://golangdocs.com/golang-download-files