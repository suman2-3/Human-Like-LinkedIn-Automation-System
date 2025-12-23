package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	AppEnv              string // dev / prod
	Email               string
	Password            string
	Headless            bool
	SlowModeMs          int
	MaxDailyConnections int
	DBPath              string
	LogLevel            string
}

var AppConfig Env

func Init() {
	_ = godotenv.Load() // ok if missing

	AppConfig = Env{
		AppEnv:              getString("APP_ENV", "dev"),
		Email:               os.Getenv("LINKEDIN_EMAIL"),
		Password:            os.Getenv("LINKEDIN_PASSWORD"),
		Headless:            getBool("HEADLESS", false),
		SlowModeMs:          getInt("SLOW_MODE_MS", 800),
		MaxDailyConnections: getInt("MAX_DAILY_CONNECTIONS", 15),
		DBPath:              getString("DB_PATH", "database/state.db"),
		LogLevel:            getString("LOG_LEVEL", "info"),
	}

	if AppConfig.Email == "" || AppConfig.Password == "" {
		log.Fatal("LINKEDIN_EMAIL or LINKEDIN_PASSWORD missing")
	}

	log.Println("🔧 Environment:", AppConfig.AppEnv)
}

func getInt(key string, fallback int) int {
	if val := os.Getenv(key); val != "" {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return fallback
}

func getBool(key string, fallback bool) bool {
	if val := os.Getenv(key); val != "" {
		return val == "true"
	}
	return fallback
}

func getString(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
