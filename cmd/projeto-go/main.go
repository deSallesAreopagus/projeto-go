package main

import (
	"log"
	"projeto-go/pkg/database"
	"projeto-go/pkg/router"

	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		log.Fatalf("Error loading .env file")
	}
	db := database.SetupDBConnection()
	defer database.CloseDBConnection(db)

	r := router.SetupRouter(db)

	r.Run(":8550")
}
