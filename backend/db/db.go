package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	 _ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connector() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := os.Getenv("DATABASE_URL")
	db , err := sql.Open("postgres", connStr)
	if err != nil { 
		log.Fatal("😢error occured in loading database",err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("😢error occured in pinging database",err)
	}
	fmt.Println("🚀Connected to database")
	DB = db
}
