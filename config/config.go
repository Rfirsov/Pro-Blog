package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		AppEnv string
		Port   string
	}

	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		SslMode  string
	}
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

	// Server config
	Cfg.Server.AppEnv = getEnv("APP_ENV", "development")
	Cfg.Server.Port = getEnv("PORT", "8080")

	// Database config
	Cfg.Database.Host = getEnv("DB_HOST", "postgres")
	Cfg.Database.Port = getEnv("DB_PORT", "5432")
	Cfg.Database.User = getEnv("DB_USER", "postgres")
	Cfg.Database.Password = getEnv("DB_PASSWORD", "")
	Cfg.Database.Name = getEnv("DB_NAME", "pro_blog_db")
	Cfg.Database.SslMode = getEnv("DB_SSLMODE", "disable")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
