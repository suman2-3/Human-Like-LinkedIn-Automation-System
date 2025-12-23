package auth

import (
	"fmt"
	"log"
	"strings"
	"time"

	"myproject/stealth"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func Login(page *rod.Page, email, password string) error {
	page.MustNavigate("https://www.linkedin.com/login")
	page.MustWaitLoad()

	stealth.HumanDelay(2000, 3000)

	emailInput := page.MustElement(`input[name="session_key"]`)
	clickAndInput(page, emailInput, email)

	stealth.HumanDelay(1200, 1800)

	passwordInput := page.MustElement(`input[name="session_password"]`)
	clickAndInput(page, passwordInput, password)

	stealth.HumanDelay(1500, 2000)

	page.MustElement(`button[type="submit"]`).MustClick()

	stealth.HumanDelay(4000, 6000)

	if detectCheckpoint(page) {
		log.Println("⚠️ LinkedIn security check detected")

		log.Println("⏸ Press ENTER after verification complete")
		fmt.Scanln()
	}

	return nil
}

func clickAndInput(page *rod.Page, el *rod.Element, text string) {
	box := el.MustShape().Box()

	x := box.X + box.Width/2
	y := box.Y + box.Height/2

	page.Mouse.MoveTo(proto.Point{X: x, Y: y})
	time.Sleep(400 * time.Millisecond)

	page.Mouse.Click(proto.InputMouseButtonLeft, 1)
	time.Sleep(500 * time.Millisecond)

	el.MustInput(text)
}

func detectCheckpoint(page *rod.Page) bool {
	url := page.MustInfo().URL
	return strings.Contains(url, "checkpoint") ||
		strings.Contains(url, "challenge")
}
