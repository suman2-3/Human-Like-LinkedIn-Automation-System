package messaging

import (
	"log"

	"github.com/go-rod/rod"
)

// This connects detector + sender + template
func RunMessagingFlow(page *rod.Page) error {
	profiles, err := GetAcceptedConnections(page)
	if err != nil {
		return err
	}

	for _, profile := range profiles {
		name := ExtractNameFromProfile(page, profile)
		msg := RenderTemplate(name)

		log.Println("Sending message to:", profile)
		SendMessage(page, profile, msg)
	}

	return nil
}
