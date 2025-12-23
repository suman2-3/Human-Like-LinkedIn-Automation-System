package stealth

import "github.com/go-rod/rod"

func ApplyFingerprint(page *rod.Page) {
	page.MustEval(`() => {
		if (navigator.webdriver !== undefined) {
			delete Object.getPrototypeOf(navigator).webdriver
		}
	}`)
}
