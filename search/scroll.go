package search

import (
	"time"

	"github.com/go-rod/rod"
)

func scrollResults(page *rod.Page) {
	for i := 0; i < 8; i++ {
		// Real mouse wheel scroll (human-like)
		page.Mouse.Scroll(0, 800, 3)

		// Wait like a human reading results
		time.Sleep(time.Duration(1200+300*i) * time.Millisecond)
	}
}
