package main

import (
	"log"

	"github.com/joho/godotenv"
	"jobhun-backend.com/database"
	"jobhun-backend.com/router"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load .env file.")
	}

	database.InitDB();

	router.Router();
}