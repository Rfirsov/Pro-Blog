package main

import (
	"log"

	"github.com/Rfirsov/Pro-Blog/routes"
)

func main() {

	router := routes.NewRouter()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Start the HTTP server on port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
