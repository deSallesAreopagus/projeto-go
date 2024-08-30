package database

import (
	"log"
	"os"
	"projeto-go/api-rest/pkg/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDBConnection() *gorm.DB {
	e := godotenv.Load("../../.env")
	if e != nil {

		e := godotenv.Load()
		if e != nil {
			log.Fatalf("Error loading .env file")
		}
	}
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
