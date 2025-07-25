package database

import (
	"fmt"
	"log"

	"github.com/Rfirsov/Pro-Blog/config"
	"github.com/Rfirsov/Pro-Blog/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		config.Cfg.Database.Host,
		config.Cfg.Database.User,
		config.Cfg.Database.Password,
		config.Cfg.Database.Name,
		config.Cfg.Database.Port,
		config.Cfg.Database.SslMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("✅ Connected to PostgreSQL via GORM")

	// Make sure uuid-ossp extension is enabled
	if err := DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		log.Fatal("failed to create uuid-ossp extension:", err)
	}

	errAutoMigrate := DB.AutoMigrate(
		&models.User{}, // Add more models here
		&models.PostStatus{},
		&models.Post{},
	)
	if errAutoMigrate != nil {
		log.Fatal("Failed to migrate database schema:", errAutoMigrate)
	}
}
