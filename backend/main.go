package main

import (
	"hotel_booking_api/controllers"
	"hotel_booking_api/database"
	"hotel_booking_api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	// Initialize database
	database.ConnectDatabase()

	// Create a Gin router
	router := gin.Default()

	//CORS Configuration:
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173"}, //change this to yours frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Public routes
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	//`CheckSession` 
	router.GET("/checksession", middleware.AuthMiddleware(), controllers.CheckSession)

	// Protected routes 
	protected := router.Group("/hotel")
	protected.Use(middleware.AuthMiddleware()) 
	{
		protected.GET("/hotellist", controllers.GetHotels)       
		protected.GET("/:id", controllers.GetHotelByID)     
		protected.POST("/book", controllers.BookHotel)      
		protected.GET("/bookings/:id", controllers.GetBookings)
		protected.GET("/review/:id", controllers.GetHotelReviews)
		protected.POST("/review", controllers.AddReview) 
	}


	router.Static("/uploads", "./uploads")



	// Start server
	router.Run(":8080")
}
