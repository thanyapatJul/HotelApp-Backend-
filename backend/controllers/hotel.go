package controllers

import (
	"hotel_booking_api/database"
	"hotel_booking_api/models"
	"net/http"
   	"fmt"
	// "gorm.io/gorm" 

	"github.com/gin-gonic/gin"
)


func GetHotelByID(c *gin.Context) {
	hotelID := c.Param("id")
	var hotel models.Hotel

	if err := database.DB.First(&hotel, hotelID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}

	c.JSON(http.StatusOK, hotel)
}


func GetHotels(c *gin.Context) {
	var hotels []models.Hotel
	database.DB.Find(&hotels)
	c.JSON(http.StatusOK, hotels)
}



func BookHotel(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userData, ok := user.(*models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user data"})
		return
	}

	var booking models.BookHotel
	if err := c.ShouldBindJSON(&booking); err != nil {
		fmt.Println("❌ Error Binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking data"})
		return
	}

	booking.CheckIn = booking.CheckIn.UTC()
	booking.CheckOut = booking.CheckOut.UTC()
	
	booking.UserID = userData.ID
	if booking.CheckIn.After(booking.CheckOut) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Check-in date must be before check-out date"})
		return
	}


	var hotel models.Hotel
	if err := database.DB.First(&hotel, booking.HotelID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}

	//Calculate total price 
	nights := int(booking.CheckOut.Sub(booking.CheckIn).Hours() / 24)
	booking.TotalPrice = float64(nights) * hotel.PricePerNight


	fmt.Println("🔍 Processed Booking Data:", booking)


	database.DB.Create(&booking)

	c.JSON(http.StatusCreated, gin.H{"message": "Hotel booked successfully", "booking": booking})
}



func GetBookings(c *gin.Context) {

	hotelID := c.Param("id")
	if hotelID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hotel ID is required"})
		return
	}


	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userData, ok := user.(*models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user data"})
		return
	}


	var bookings []models.BookHotel
	if err := database.DB.Where("hotel_id = ?", hotelID).Preload("User").Preload("Hotel").Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving bookings"})
		return
	}


	var formattedBookings []map[string]interface{}
	for _, booking := range bookings {
		formattedBookings = append(formattedBookings, map[string]interface{}{
			"id":         booking.ID,
			"userId":     booking.UserID,
			"hotelId":    booking.HotelID,
			"checkIn":    booking.CheckIn,
			"checkOut":   booking.CheckOut,
			"totalPrice": booking.TotalPrice,
			"name":       booking.Name,
			"telephone":  booking.Telephone,
			"email":      booking.Email,
			"isMine":     booking.UserID == userData.ID, 
		})
	}


	if len(formattedBookings) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No bookings found for this hotel"})
		return
	}

	c.JSON(http.StatusOK, formattedBookings)
}


func GetBookingsByHotel(c *gin.Context) {
	hotelID := c.Param("id")
	var bookings []models.BookHotel

	database.DB.Where("hotel_id = ?", hotelID).Preload("User").Preload("Hotel").Find(&bookings)
	if len(bookings) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No bookings found for this hotel"})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

func GetHotelReviews(c *gin.Context) {
	hotelID := c.Param("id")

	var reviews []models.Review
	if err := database.DB.Preload("User").Where("hotel_id = ?", hotelID).Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving reviews"})
		return
	}

	var formattedReviews []map[string]interface{}
	for _, review := range reviews {
		formattedReviews = append(formattedReviews, map[string]interface{}{
			"id":        review.ID,
			"userId":    review.UserID,
			"hotelId":   review.HotelID,
			"username":  review.User.Username,  
			"rating":    review.Rating,
			"comment":   review.Comment,
			"createdAt": review.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, formattedReviews)
}

func AddReview(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userData, ok := user.(*models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user data"})
		return
	}

	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review data"})
		return
	}

	review.UserID = userData.ID


	if review.Rating < 1 || review.Rating > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rating must be between 1 and 5"})
		return
	}

	var hotel models.Hotel
	if err := database.DB.First(&hotel, review.HotelID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}

	database.DB.Create(&review)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Review added successfully",
		"review":  review,
	})
}
