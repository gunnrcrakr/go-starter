package bootstrap

import (
	"go-starter/internal/database"
	"log"

	"github.com/joho/godotenv"
)

func Starter() {

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Connect database
	database.InitDB()
}
