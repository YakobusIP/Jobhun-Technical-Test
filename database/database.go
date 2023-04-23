package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB_USER string
	DB_PASS string
	DB_NAME string
	DATABASE *sql.DB
)

func getDBCredentials() {
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_NAME = os.Getenv("DB_NAME")
}

func InitDB() {
	getDBCredentials()

	connString := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", DB_USER, DB_PASS, DB_NAME)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database successfully connected");
	}

	DATABASE = db
}