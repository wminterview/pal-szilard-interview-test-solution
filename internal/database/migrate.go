package database

import (
	"library/internal/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	// Automatikus migráció a Book és Borrowing modellekhez
	if err := db.AutoMigrate(&models.Book{}, &models.Borrowing{}); err != nil {
		return err
	}
	return nil
}
