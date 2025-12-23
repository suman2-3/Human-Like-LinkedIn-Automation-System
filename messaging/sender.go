package messaging

import (
	"time"

	"github.com/go-rod/rod"
)

func SendMessage(page *rod.Page, profileURL string, msg string) error {
	// Navigate to profile
	page.MustNavigate(profileURL)
	page.MustWaitLoad()
	time.Sleep(2 * time.Second)

	// SAFE: Check Message button exists
	messageBtn, err := page.Timeout(5 * time.Second).
		Element(`button:has-text("Message")`)
	if err != nil {
		// Messaging not allowed / already messaged
		return nil
	}

	messageBtn.MustClick()
	time.Sleep(2 * time.Second)

	// Type message
	input := page.MustElement(`div[role="textbox"]`)
	input.MustInput(msg)
	time.Sleep(800 * time.Millisecond)

	// Click Send
	page.MustElement(`button:has-text("Send")`).MustClick()
	time.Sleep(2 * time.Second)

	return nil
}
