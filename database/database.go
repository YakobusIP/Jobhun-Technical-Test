package database

import (
	"database/sql"
	"fmt"
)

const (
	DB_USER = "root"
	DB_NAME = "jobhun_db"
)

var (
	DATABASE *sql.DB
)

func InitDB() {
	db, err := sql.Open("mysql", "root:Yakobus-13520104@tcp(localhost:3306)/jobhun_db?parseTime=true")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected successfully");
	}

	DATABASE = db
}