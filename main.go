package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"jobhun-backend.com/database"
	"jobhun-backend.com/router"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Failed to load .env file")
	}
	
	database.InitDB();

	router.Router();
}