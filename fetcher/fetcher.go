package fetcher

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func ReadHtml(urlString string) (string, error) {
	resp, err := http.Get(urlString)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%d %s", resp.StatusCode, resp.Status)
	}

	htmlBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(htmlBytes), nil
}

func DownloadAtPath(urlString, atPath string) error {
	resp, err := http.Get(urlString)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(atPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
