package search

import (
	"strings"

	"github.com/go-rod/rod"
)

func extractProfileLinks(page *rod.Page) []string {
	links := []string{}

	elems := page.MustElements(`a[href*="/in/"]`)
	seen := map[string]bool{}

	for _, el := range elems {
		href, err := el.Attribute("href")
		if err == nil && href != nil {
			link := *href

			if idx := strings.Index(link, "?"); idx != -1 {
				link = link[:idx]
			}

			if !seen[link] {
				seen[link] = true
				links = append(links, link)
			}
		}
	}

	return links
}
