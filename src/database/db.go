package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB = nil
)

func ConnectDB() error {
	log.Println("Connecting to DB")

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPortString := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")

	dbPort, err := strconv.Atoi(dbPortString)
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Session{})
	db.AutoMigrate()

	DB = db
	return nil
}
