package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var con *sql.DB

func Connect() *sql.DB { // função Connect com instãncia de sql.DB

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	USER := os.Getenv("DBUSER")
	PASSWORD := os.Getenv("DBPASSWORD")
	HOST := os.Getenv("DBHOST")
	PORT := os.Getenv("DBPORT")
	DATABASE := os.Getenv("DBDATABASE")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", USER, PASSWORD, HOST, PORT, DATABASE)
	fmt.Println(dataSourceName)
	db, err := sql.Open("mysql", dataSourceName)
	fmt.Println(db, "\n", err)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database");
	con = db
	return db
}
