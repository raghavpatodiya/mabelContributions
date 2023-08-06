package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var createPinBoardsTableQuery = `CREATE TABLE IF NOT EXISTS pin_boards (
	board_id TEXT PRIMARY KEY, 
	username TEXT NOT NULL, 
	board_name TEXT NOT NULL,
	is_public BOOLEAN NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,
	is_updated BOOLEAN,
	updated_at TIMESTAMP WITH TIME ZONE,
	is_deleted BOOLEAN,
	deleted_at TIMESTAMP WITH TIME ZONE,
	likes_count Numeric NOT NULL,
	shares_count Numeric NOT NULL,

	CONSTRAINT unique_username_boardname UNIQUE (username, board_name)
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

	if err := db.Exec(createPinBoardsTableQuery).Error; err != nil {
		fmt.Println("error creating pin boards table: ", err)
		return
	}

	fmt.Println("----successfully created table pin_boards----")
}
