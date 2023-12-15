package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	

	dbURL := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbUser, dbName, dbPassword)
	
	db, err := gorm.Open("postgres", dbURL)


	if err != nil {
		log.Fatal(err)
	}
	return db
}
