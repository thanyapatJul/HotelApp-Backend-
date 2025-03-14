package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	// MySQL connection string
	dsn := "hotel_user:o863454681@tcp(127.0.0.1:3306)/hotel_db?charset=utf8mb4&parseTime=True&loc=Local"
	
	// Open connection
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	// Migrate tables
	db.AutoMigrate(&User{}, &Booking{})

	DB = db
	fmt.Println("Connected to the database!")
}
