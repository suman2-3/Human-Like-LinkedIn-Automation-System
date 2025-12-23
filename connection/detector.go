package connection

import (
	"strings"

	"github.com/go-rod/rod"
)

func HasConnectButton(page *rod.Page) bool {
	btns, _ := page.Elements("button")

	for _, btn := range btns {
		label, _ := btn.Attribute("aria-label")
		if label == nil {
			continue
		}

		text := strings.ToLower(*label)
		if strings.Contains(text, "invite") ||
			strings.Contains(text, "connect") ||
			strings.Contains(text, "add") {
			return true
		}
	}

	return false
}
