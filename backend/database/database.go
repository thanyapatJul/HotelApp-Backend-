package database

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/datatypes"

	"hotel_booking_api/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	// Replace with your MySQL password
	dsn := "root:PASSWORD@tcp(127.0.0.1:3306)/hotel_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	fmt.Println("✅ Connected to MySQL successfully!")


	DB.AutoMigrate(&models.User{}, &models.Hotel{}, &models.BookHotel{}, &models.Review{})


	SeedDatabase(DB)
}


func toJSON(data interface{}) datatypes.JSON {
	jsonData, _ := json.Marshal(data)
	return datatypes.JSON(jsonData)
}

func SeedDatabase(db *gorm.DB) {
	fmt.Println("🔄 Resetting database...")
	db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	db.Exec("TRUNCATE TABLE reviews;")
	db.Exec("TRUNCATE TABLE book_hotels;")
	db.Exec("TRUNCATE TABLE hotels;")

	db.Exec("SET FOREIGN_KEY_CHECKS=1;")
	fmt.Println("✅ Database reset!")


	users := []models.User{
		{Username: "Premeadmin", Email: "preme@example.com", Password: "123456"},
		{Username: "Alice", Email: "alice@example.com", Password: "password123"},
		{Username: "Bob", Email: "bob@example.com", Password: "password456"},
	}
	db.Create(&users)


	var createdUsers []models.User
	db.Find(&createdUsers)

	if len(createdUsers) < 2 {
		log.Fatal("❌ Not enough users for review seeding!")
	}

	user1ID := createdUsers[0].ID
	user2ID := createdUsers[1].ID


	hotels := []models.Hotel{
		{
			Name:          "Grand Hotel",
			Location:      "Bangkok",
			PricePerNight: 2000,
			Rooms:         3,
			Rating:        4.5,
			Images:        toJSON([]string{"uploads/image1.jpg", "uploads/image2.jpg"}), 
		},
		{
			Name:          "Luxury Resort",
			Location:      "Phuket",
			PricePerNight: 3500,
			Rooms:         3,
			Rating:        5.0,
			Images:        toJSON([]string{"uploads/image2.jpg", "uploads/image3.jpg"}), 
		},
		{
			Name:          "Le Resort",
			Location:      "Rayong",
			PricePerNight: 3500,
			Rooms:         3,
			Rating:        5.0,
			Images:        toJSON([]string{"uploads/image3.jpg", "uploads/image3.jpg"}), 
		},
	}
	db.Create(&hotels)


	var createdHotels []models.Hotel
	db.Find(&createdHotels)

	if len(createdHotels) < 2 {
		log.Fatal("❌ Not enough hotels for review seeding!")
	}

	hotel1ID := createdHotels[0].ID
	hotel2ID := createdHotels[1].ID


	reviews := []models.Review{
		{UserID: user1ID, HotelID: hotel1ID, Rating: 4.5, Comment: "Amazing experience!"},
		{UserID: user2ID, HotelID: hotel1ID, Rating: 3.8, Comment: "Pretty good but expensive."},
		{UserID: user1ID, HotelID: hotel2ID, Rating: 5.0, Comment: "Best hotel ever!"},
	}
	db.Create(&reviews)

	fmt.Println("✅ Seeding complete!")
}
