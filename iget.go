package main

import (
	"iget/downloader"
	"os"
)

func main() {
	url := "https://velog.io/@kineo2k/golang-hasty-abstractions"
	savePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dl := downloader.New(savePath)
	err = dl.Get(url)
	if err != nil {
		panic(err)
	}
}
