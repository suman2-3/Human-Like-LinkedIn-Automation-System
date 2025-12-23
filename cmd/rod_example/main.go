package main

import (
	"log"

	"myproject/auth"
	"myproject/config"
	"myproject/connection"
	"myproject/search"
	"myproject/stealth"
	"myproject/utils"
)

func main() {
	config.Init()

	browser, page := utils.NewBrowser()
	defer browser.MustClose()

	stealth.ApplyFingerprint(page)

	err := auth.Login(
		page,
		config.AppConfig.Email,
		config.AppConfig.Password,
	)

	if err != nil {
		log.Println("⚠️ Login interrupted:", err.Error())

	}

	profiles, err := search.PeopleSearch(page, "software engineer")
	if err != nil {
		log.Println("Search error:", err)
		return
	}

	limiter := connection.NewLimiter(config.AppConfig.MaxDailyConnections)

	for _, profile := range profiles {
		if !limiter.Allow() {
			log.Println("🚫 Daily limit reached")
			break
		}

		if err := connection.SendConnection(page, profile); err != nil {
			log.Println("Error:", err)
		}

		connection.Cooldown()
	}

	select {}
}
