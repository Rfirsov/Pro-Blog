package seeders

import (
	"log"

	"gorm.io/gorm"
)

// SeedAll runs all seeders
func SeedAll(db *gorm.DB) {
	log.Println("🔁 Seeding database...")

	if err := SeedPostStatuses(db); err != nil {
		log.Fatalf("❌ Failed to seed post statuses: %v", err)
	}

	log.Println("✅ Database seeded successfully.")
}
