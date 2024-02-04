package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB = nil
)

func ConnectDB() error {
	// TODO: ADD CONNECTION DETAILS
	dsn := fmt.Sprintf("")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
