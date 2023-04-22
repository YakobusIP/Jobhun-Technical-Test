package main

import (
	"fmt"

	"jobhun-backend.com/database"
	"jobhun-backend.com/router"
)

func main() {
	fmt.Println("Hello world");
	database.InitDB();

	router.Router();
}