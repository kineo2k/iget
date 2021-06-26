package downloader

import (
	"fmt"
	. "iget/collections"
	_ "iget/collections"
	"iget/fetcher"
	"iget/savepath"
	"net/url"
	"sync"
)

type Downloader struct {
	urlString string
	savePath  *savepath.SavePath
}

func New(urlString string) *Downloader {
	return &Downloader{
		urlString: urlString,
		savePath:  savepath.New(domainFromUrl(urlString)),
	}
}

func (d *Downloader) Get() error {
	fmt.Println(fmt.Sprintf("Loading HTML from a %s.", d.urlString))
	html, err := fetcher.ReadHtml(d.urlString)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Parse image URLs from HTML."))
	urls := fetcher.ParseImgUrls(html)
	numOfImages := urls.Len()

	if numOfImages == 0 {
		fmt.Println("Image URL not found.")
		return nil
	} else {
		fmt.Println(fmt.Sprintf("Found %d image URLs.", numOfImages))

		err := d.savePath.Create()
		if err != nil {
			return err
		}
	}

	d.downloadConcurrency(urls)

	return nil
}

func (d *Downloader) downloadConcurrency(urls *Set) {
	wg := sync.WaitGroup{}
	wg.Add(urls.Len())

	download := func(urlString, path string) {
		defer wg.Done()

		err := fetcher.DownloadAtPath(urlString, path)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(fmt.Sprintf("I got it [%s]", path))
		}
	}

	for _, urlString := range urls.Entries() {
		go download(urlString, d.savePath.WithUrl(urlString))
	}
	wg.Wait()
}

func domainFromUrl(urlString string) string {
	u, err := url.Parse(urlString)
	if err != nil {
		fmt.Println(err)
	}

	return u.Host
}
