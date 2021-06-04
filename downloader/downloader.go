package downloader

import (
	"fmt"
	. "iget/collections"
	_ "iget/collections"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

var imgElementRegex = regexp.MustCompile("<img.*?src=\"(.*?)\"[^\\>]+>")
var urlRegex = regexp.MustCompile("(http:\\/\\/|https:\\/\\/)[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*\\.(jpg|png|gif|svg|webp))?")

type Downloader struct {
	savePath string
}

func New(savePath string) *Downloader {
	return &Downloader{savePath}
}

func (d *Downloader) Get(urlString string) error {
	fmt.Println(fmt.Sprintf("Loading HTML from a %s.", urlString))
	html, err := d.readHtml(urlString)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Parse image URLs from HTML."))
	urls := d.parseImgUrls(html)
	numOfImages := urls.Len()
	domain := d.domainFromUrl(urlString)

	if numOfImages == 0 {
		fmt.Println("Image URL not found.")
		return nil
	} else {
		fmt.Println(fmt.Sprintf("Found %d image URLs.", numOfImages))

		err := d.createDownloadDirectory(domain)
		if err != nil {
			return err
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(numOfImages)

	download := func(url, path string) {
		defer wg.Done()

		err := d.downloadAtPath(url, path)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(fmt.Sprintf("I got it [%s]", path))
		}
	}

	for _, urlString := range urls.Entries() {
		go download(urlString, d.downloadPathWithURL(urlString, domain))
	}
	wg.Wait()

	return nil
}

func (d *Downloader) readHtml(urlString string) (string, error) {
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

func (d *Downloader) parseImgUrls(html string) *Set {
	imgElements := imgElementRegex.FindAllString(html, -1)
	set := NewSet()
	for _, elem := range imgElements {
		urls := urlRegex.FindAllString(elem, -1)
		for _, url := range urls {
			fmt.Println(fmt.Sprintf("Found [%s]", url))
			set.Add(url)
		}
	}

	return set
}

func (d *Downloader) domainFromUrl(urlString string) string {
	u, err := url.Parse(urlString)
	if err != nil {
		log.Fatal(err)
	}

	return u.Host
}

func (d *Downloader) downloadPathWithURL(urlString, domain string) string {
	fn := filepath.Base(urlString)

	return fmt.Sprintf("%s%c%s%c%s", d.savePath, os.PathSeparator, domain, os.PathSeparator, fn)
}

func (d *Downloader) createDownloadDirectory(domain string) error {
	return os.Mkdir(fmt.Sprintf("%s%c%s", d.savePath, os.PathSeparator, domain), os.ModePerm)
}

func (d *Downloader) downloadAtPath(urlString, atPath string) error {
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
