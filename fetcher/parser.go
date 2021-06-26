package fetcher

import (
	"fmt"
	"iget/collections"
	"regexp"
)

var imgElementRegex = regexp.MustCompile("<img.*?src=\"(.*?)\"[^\\>]+>")
var urlRegex = regexp.MustCompile("(http:\\/\\/|https:\\/\\/)[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*\\.(jpg|jpeg|png|gif|svg|webp))?")

func ParseImgUrls(html string) *collections.Set {
	imgElements := imgElementRegex.FindAllString(html, -1)
	set := collections.NewSet()
	for _, elem := range imgElements {
		urls := urlRegex.FindAllString(elem, -1)
		for _, urlString := range urls {
			fmt.Println(fmt.Sprintf("Found [%s]", urlString))
			set.Add(urlString)
		}
	}

	return set
}
