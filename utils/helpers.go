package utils

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// Human-like random delay
func HumanDelay(minMs, maxMs int) {
	if maxMs <= minMs {
		time.Sleep(time.Duration(minMs) * time.Millisecond)
		return
	}
	delay := minMs + rand.Intn(maxMs-minMs)
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

// Random hover to simulate curiosity
func RandomHover(page *rod.Page) {
	els, err := page.Elements("button, a, div")
	if err != nil || len(els) == 0 {
		return
	}
	_ = els[rand.Intn(len(els))].Hover()
}

// Human-like small scroll
func SmallScroll(page *rod.Page) {
	page.Mouse.Scroll(
		0,
		float64(rand.Intn(200)+100), // float64 FIX
		1,
	)
}
