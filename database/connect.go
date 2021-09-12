package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/Vardan1995/fiber-crud/model"
)

// ConnectDB connect to db
func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	fmt.Println(DB)
  if err != nil {
    panic("failed to connect database")
  }
	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate( &model.User{})

	fmt.Println("Database Migrated")
}