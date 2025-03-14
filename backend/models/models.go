package models

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"gorm.io/datatypes"
)


type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Hotel struct {
	ID            uint           `gorm:"primaryKey"`
	Name          string         `gorm:"type:varchar(255);not null"` 
	Location      string         `gorm:"type:varchar(255);not null"`
	PricePerNight float64        `gorm:"not null"`
	Rooms         int            `gorm:"not null"`
	Rating        float64        `gorm:"not null;default:0"`
	Images        datatypes.JSON `gorm:"type:json"` 
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	
}


type BookHotel struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `json:"userId" gorm:"not null"` 
	User       User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	HotelID    uint      `json:"hotelId"`
	Hotel      Hotel     `gorm:"foreignKey:HotelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name       string    `json:"name"`
	Telephone  string    `json:"telephone"`
	Email      string    `json:"email"`
	CheckIn    time.Time `json:"checkIn"`
	CheckOut   time.Time `json:"checkOut"`

	TotalPrice float64   `json:"totalPrice"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Review struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID"`
	HotelID   uint      `gorm:"not null"`
	Hotel     Hotel     `gorm:"foreignKey:HotelID"`
	Rating    float64   `gorm:"not null"`
	Comment   string    `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
