package utils

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// SafeClick finds element fresh every time and clicks safely
func SafeClick(page *rod.Page, selector string) bool {
	el, err := page.Timeout(5 * time.Second).Element(selector)
	if err != nil {
		return false
	}

	_ = el.ScrollIntoView()

	if err := el.WaitEnabled(); err != nil {
		return false
	}

	time.Sleep(time.Duration(300+rand.Intn(700)) * time.Millisecond)

	if err := el.Click(proto.InputMouseButtonLeft, 1); err != nil {
		return false
	}

	return true
}
