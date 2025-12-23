package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func RandomScroll(page *rod.Page) {
	steps := rand.Intn(3) + 2

	for i := 0; i < steps; i++ {
		page.Mouse.Scroll(
			0,
			float64(rand.Intn(400)+200),
			1, // REQUIRED by rod v0.116
		)

		time.Sleep(time.Duration(rand.Intn(800)+400) * time.Millisecond)
	}
}
