package connection

import (
	"log"
	"time"

	"myproject/utils"

	"github.com/go-rod/rod"
)

func SendConnection(page *rod.Page, profileURL string) error {
	log.Println("Visiting:", profileURL)

	page.MustNavigate(profileURL)
	page.MustWaitLoad()
	page.MustWaitIdle()
	time.Sleep(2 * time.Second)

	if !HasConnectButton(page) {
		log.Println("No Add/Connect available, skipping")
		return nil
	}

	// SAFE CLICK using aria-label
	clicked := false
	btns, _ := page.Elements("button")

	for _, btn := range btns {
		label, _ := btn.Attribute("aria-label")
		if label == nil {
			continue
		}

		if utils.SafeClick(page, `button[aria-label="`+*label+`"]`) {
			log.Println("✅ Clicked Add/Connect")
			clicked = true
			break
		}
	}

	if !clicked {
		log.Println("Failed to click Add safely")
		return nil
	}

	time.Sleep(2 * time.Second)

	// MODAL HANDLING (THIS IS CRITICAL)
	if utils.SafeClick(page, `div[role="dialog"] button:has-text("Send without a note")`) {
		log.Println("🎉 Connection sent")
		time.Sleep(3 * time.Second)
		return nil
	}

	log.Println("Send button not available (restricted)")
	return nil
}
