package sanitize

import (
	"net/url"

	p "github.com/PuerkitoBio/purell"
)

func Url(rawUrl string) (string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	return p.NormalizeURL(u, p.FlagsSafe|p.FlagRemoveDuplicateSlashes), nil
}
