package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//set GO111MODULE=on
//go run cmd/scripts/create_feed_shares_table/main.go

var createFeedSharesTableQuery = `CREATE TABLE IF NOT EXISTS feed_shares(
	username TEXT NOT NULL, 
	feed_id TEXT NOT NULL, 

	CONSTRAINT pk_feed_share PRIMARY KEY (username, feed_id)
	);`

func main() {
	// Db URL from env
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("error loading environment file: ", err)
	}

	dbUrl := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	DB_URL := "postgres://" + dbName + ":" + dbPassword + "@" + dbUrl

	db, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
	if err != nil {
		fmt.Println("error connecting to database: ", err)
		return
	}

	if err := db.Exec(createFeedSharesTableQuery).Error; err != nil {
		fmt.Println("error creating feed shares table: ", err)
		return
	}

	fmt.Println("----successfully created table feed_shares----")
}
