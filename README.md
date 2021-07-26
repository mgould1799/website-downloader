# Website Downloader

## Description

This is a simple program to download sites and save them to a directoy. This is to showcase a simple implementation of the pipeline concurrency pattern in golang.

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

#### Flags that can be passed

`--urls`
Urls is a list of comma separated urls to pull down and save

example run 
```
./WebsiteDownloader --urls=https://google.com,https://facebook.com
````

`--storageLocation`
example run 
```
./WebsiteDownloader --storageLocation=./temp/test
```

Full example
```
./WebsiteDownloader --urls=https://google.com,https://facebook.com,https://instagram.com,https://stackoverflow.com,https://microsoft.com,https://github.com,gitlab.com --storageLocation=./temp/test
```

## Acknowledgments

Inspiration, code snippets, etc.
* https://golangdocs.com/golang-download-files
* https://stackoverflow.com/questions/55203251/limiting-number-of-go-routines-running
* https://medium.com/@deckarep/gos-extended-concurrency-semaphores-part-1-5eeabfa351ce



## License
 
The MIT License (MIT)

Copyright (c) 2021 Meagan Gould

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.