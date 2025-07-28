package seeders

import (
	"log"

	"github.com/Rfirsov/Pro-Blog/internal/models"
	"gorm.io/gorm"
)

func SeedPostTags(db *gorm.DB) error {
	postTags := []string{"go", "backend", "programming", "tutorial"}

	for _, name := range postTags {
		postTag := models.Tag{Name: name}
		if err := db.FirstOrCreate(&postTag, models.Tag{Name: name}).Error; err != nil {
			log.Printf("Failed to seed tag '%s': %v", name, err)
			return err
		}
	}

	return nil
}
