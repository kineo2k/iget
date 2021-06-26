package downloader

import (
	"testing"
)

func TestGet(t *testing.T) {
	urlString := "https://velog.io/@kineo2k/golang-hasty-abstractions"

	dl := New(urlString)
	err := dl.Get()
	if err != nil {
		t.Fail()
	}
}
