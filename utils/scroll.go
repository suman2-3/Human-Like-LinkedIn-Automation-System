package utils

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func RandomScroll(page *rod.Page) {
	steps := rand.Intn(5) + 3
	for i := 0; i < steps; i++ {
		page.MustEval(`() => window.scrollBy(0, Math.floor(Math.random()*400)+200)`)
		time.Sleep(time.Duration(500+rand.Intn(1000)) * time.Millisecond)
	}
}
