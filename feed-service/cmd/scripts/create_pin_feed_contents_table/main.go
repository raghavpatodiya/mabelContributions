package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var createPinFeedContentsTableQuery = `CREATE TABLE IF NOT EXISTS pin_feed_contents (
	feed_id TEXT NOT NULL,
	board_id TEXT NOT NULL, 

	CONSTRAINT unique_feed_content_pin_board UNIQUE (feed_id, board_id)
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

	if err := db.Exec(createPinFeedContentsTableQuery).Error; err != nil {
		fmt.Println("error creating pin feed contents table: ", err)
		return
	}

	fmt.Println("----successfully created table pin_feed_contents----")
}
