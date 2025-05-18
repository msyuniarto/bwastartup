package database

import (
	"bwastartup/campaign"
	"bwastartup/transaction"
	"bwastartup/user"

	"gorm.io/gorm"
)

// AutoMigrate melakukan migrasi skema database berdasarkan model yang diberikan.
func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&campaign.Campaign{},
		&campaign.CampaignImage{},
		&transaction.Transaction{},
		&user.User{},
	)
	if err != nil {
		panic("failed to migrate database")
	}
	println("Database migration completed successfully.")
}
