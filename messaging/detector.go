package messaging

import (
	"strings"
	"time"

	"github.com/go-rod/rod"
)

func GetAcceptedConnections(page *rod.Page) ([]string, error) {
	page.MustNavigate("https://www.linkedin.com/mynetwork/invite-connect/connections/")
	page.MustWaitLoad()
	time.Sleep(2 * time.Second)

	var profiles []string
	seen := make(map[string]bool)

	//FIX
	links := page.MustElements(`a[href*="/in/"]`)
	for _, el := range links {
		href, err := el.Attribute("href")
		if err != nil || href == nil {
			continue
		}

		url := cleanProfileURL(*href)
		if !seen[url] {
			seen[url] = true
			profiles = append(profiles, url)
		}
	}

	return profiles, nil
}

// Name extraction (UNCHANGED)
func ExtractNameFromProfile(page *rod.Page, profileURL string) string {
	page.MustNavigate(profileURL)
	page.MustWaitLoad()
	time.Sleep(1 * time.Second)

	el, err := page.Timeout(3 * time.Second).
		Element(`h1`)
	if err != nil {
		return "there"
	}

	return el.MustText()
}

// Helper: clean LinkedIn profile URLs
func cleanProfileURL(url string) string {
	if strings.Contains(url, "?") {
		return strings.Split(url, "?")[0]
	}
	return url
}
