package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// HoverElement simulates human-like mouse hovering
func HoverElement(page *rod.Page, el *rod.Element) {
	box := el.MustShape().Box()

	// Random point inside element
	x := box.X + rand.Float64()*box.Width
	y := box.Y + rand.Float64()*box.Height

	page.Mouse.MoveTo(proto.Point{X: x, Y: y})
	time.Sleep(time.Duration(rand.Intn(500)+300) * time.Millisecond)
}
