package database

import (
	"log"
	"os"
	"projeto-go/api-rest/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDBConnection() *gorm.DB {

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&models.Book{})
	DB = db
	return db
}

func CloseDBConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatal("Failed to close DB connection: ", err)
	}
	dbSQL.Close()
}
