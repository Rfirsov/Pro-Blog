package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv     string
	ServerPort string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var Cfg Config

func LoadConfig() {
	envFile := ".env.dev" // default

	// Detect which .env file to load based on APP_ENV or fallback
	appEnv := os.Getenv("APP_ENV")
	switch appEnv {
	case "development":
		envFile = ".env.dev"
	case "production":
		envFile = ".env.prod"
	}

	// Load the env file, fallback if missing
	if err := godotenv.Load(".env.dev"); err != nil {
		log.Printf("No %s file found, using environment variables\n", envFile)
	}

	// Now load config values
	Cfg = Config{
		AppEnv:     getEnv("APP_ENV", "development"),
		ServerPort: getEnv("PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "postgres"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "pro_blog_db"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
