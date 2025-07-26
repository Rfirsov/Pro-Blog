package seeders

import (
	"log"

	"github.com/Rfirsov/Pro-Blog/internal/models"
	"gorm.io/gorm"
)

func SeedPostStatuses(db *gorm.DB) error {
	statuses := []models.PostStatus{
		{Value: "draft", Label: "Draft"},
		{Value: "published", Label: "Published"},
		{Value: "archived", Label: "Archived"},
	}

	for _, status := range statuses {
		if err := db.FirstOrCreate(&status, models.PostStatus{Value: status.Value}).Error; err != nil {
			log.Printf("Failed to seed post status: %s - %v", status.Value, err)
			return err
		}
	}

	return nil
}
