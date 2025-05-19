package database

import (
	"fmt"
	"log"

	// Sesuaikan dengan driver database Anda
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// sqlite version
func SetupDatabase1() *gorm.DB {
	// Buat koneksi ke SQLite (akan membuat file jika belum ada)
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to SQLite database:", err)
		return nil
	}

	fmt.Println("Connection to SQLite database is good")
	return db
}
