package main

import (
	"iget/downloader"
	"testing"
)

func TestIGet(t *testing.T) {
	urlString := "https://velog.io/@kineo2k/golang-hasty-abstractions"

	dl := downloader.New(urlString)
	err := dl.Get()
	if err != nil {
		t.Fail()
	}
}
