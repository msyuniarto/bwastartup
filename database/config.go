package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql" // Sesuaikan dengan driver database Anda
	"gorm.io/gorm"
)

// SetupDatabase menginisialisasi koneksi database GORM.
func SetupDatabase() *gorm.DB {
	// Load file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil // Atau panic, tergantung bagaimana Anda ingin menangani error
	}

	// Ambil nilai dari .env
	ipAddress := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, ipAddress, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		return nil // Atau panic
	}

	fmt.Println("Connection to database is good from config file")
	return db
}

// sqlite version
// func SetupDatabase() *gorm.DB {
// 	// Buat koneksi ke SQLite (akan membuat file jika belum ada)
// 	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Failed to connect to SQLite database:", err)
// 		return nil
// 	}

// 	fmt.Println("Connection to SQLite database is good")
// 	return db
// }
