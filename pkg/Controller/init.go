package controller

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	// Open a connection to the MySQL database
	var err error
	db, err = sql.Open("mysql", "root:admin@tcp(localhost:3306)/parking")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected successfully")
}
