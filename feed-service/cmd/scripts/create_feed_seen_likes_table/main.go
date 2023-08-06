package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var query = `CREATE TABLE IF NOT EXISTS feed_seen_likes(
	username TEXT NOT NULL, 
	feed_id TEXT NOT NULL, 
	liked BOOLEAN,

	CONSTRAINT pk_feed_seen_likes PRIMARY KEY (username, feed_id)
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

	if err := db.Exec(query).Error; err != nil {
		fmt.Println("error creating feed seen likes table: ", err)
		return
	}

	fmt.Println("----successfully created table feed_seen_likes----")
}
