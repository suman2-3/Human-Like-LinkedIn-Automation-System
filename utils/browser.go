package utils

import (
	"myproject/config"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func NewBrowser() (*rod.Browser, *rod.Page) {
	l := launcher.New().
		Headless(config.AppConfig.Headless).
		NoSandbox(true).
		Leakless(false)

	url := l.MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()

	page := browser.MustPage()

	return browser, page
}
