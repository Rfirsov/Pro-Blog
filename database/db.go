package database

import (
	"fmt"
	"log"
	"os/user"

	"github.com/Rfirsov/Pro-Blog/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		config.Cfg.DBHost,
		config.Cfg.DBUser,
		config.Cfg.DBPassword,
		config.Cfg.DBName,
		config.Cfg.DBPort,
		config.Cfg.DBSslMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("✅ Connected to PostgreSQL via GORM")

	errAutoMigrate := DB.AutoMigrate(
		&user.User{}, // Add more models here
	)
	if errAutoMigrate != nil {
		log.Fatal("Failed to migrate database schema:", errAutoMigrate)
	}
}
