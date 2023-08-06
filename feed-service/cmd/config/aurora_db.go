package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Db URL from env
	dbUrl := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	DB_URL := "postgres://" + dbName + ":" + dbPassword + "@" + dbUrl

	// Create a new DB connection
	db, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	psdb, err := db.DB()
	if err != nil {
		panic(err)
	}
	// Config connection pool
	psdb.SetMaxIdleConns(10)
	psdb.SetMaxOpenConns(100)
	// Set DB connection as global
	DB = db
}
