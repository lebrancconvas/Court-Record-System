package db

import (
	"lebrancconvas/courtrecord/config"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// dsn := "host=localhost user=courtrecord password=P@ssw0rd dbname=courtrecord port=5402"
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port
	dial := postgres.Open(dsn)
	db, err = gorm.Open(dial, config.Config)
	if err != nil {
		panic(err)
	}
}

func GetDB() (*gorm.DB) {
	return db
}