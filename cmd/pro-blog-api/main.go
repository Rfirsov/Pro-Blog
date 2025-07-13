// main.go

// @title           ProBlog API
// @version         1.0
// @description     API for a blog platform built with Go, Gin, GORM, JWT
// @termsOfService  http://swagger.io/terms/

// @contact.name   Roman
// @contact.email  your.email@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

package main

import (
	"log"

	"github.com/Rfirsov/Pro-Blog/config"
	"github.com/Rfirsov/Pro-Blog/database"
	_ "github.com/Rfirsov/Pro-Blog/docs"
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
