package search

import (
	"net/url"
	"time"

	"github.com/go-rod/rod"
)

func PeopleSearch(page *rod.Page, keyword string) ([]string, error) {
	searchURL := buildPeopleSearchURL(keyword)

	page.MustNavigate(searchURL)
	page.MustWaitLoad()
	time.Sleep(3 * time.Second)

	scrollResults(page)

	profiles := extractProfileLinks(page)

	return profiles, nil
}

func buildPeopleSearchURL(keyword string) string {
	q := url.Values{}
	q.Set("keywords", keyword)

	return "https://www.linkedin.com/search/results/people/?" + q.Encode()
}
