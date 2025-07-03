package main

import (
	"log"

	"github.com/Rfirsov/Pro-Blog/config"
	"github.com/Rfirsov/Pro-Blog/database"
	"github.com/Rfirsov/Pro-Blog/routes"
)

func main() {
	config.LoadConfig()
	database.InitDB()

	router := routes.NewRouter()
	if config.Cfg.Server.AppEnv == "development" {
		router.SetTrustedProxies([]string{"127.0.0.1"})
	}

	if err := router.Run(":" + config.Cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
