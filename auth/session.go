package auth

import (
	"encoding/json"
	"os"

	"github.com/go-rod/rod"
)

func SaveSession(page *rod.Page, path string) error {
	cookies, err := page.Cookies(nil)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cookies, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
