package main

import (
	"github.com/gemdivk/LUMERA-SPA/api-gateway/booking"
	"github.com/gemdivk/LUMERA-SPA/api-gateway/review"
	"github.com/gemdivk/LUMERA-SPA/api-gateway/user"
	"github.com/gin-gonic/gin"
)

func main() {
	review.InitGRPCClient()
	user.InitGRPCClient()
	booking.InitGRPCClient()

	router := gin.Default()

	rg := router.Group("/reviews")
	{
		rg.POST("/", review.CreateReview)
		rg.GET("/:id", review.GetReview)
		rg.PUT("/:id", review.UpdateReview)
		rg.DELETE("/:id", review.DeleteReview)
	}
	router.GET("/salons/:salon_id/reviews", review.ListReviews)
	userGroup := router.Group("/users")
	{
		userGroup.POST("/register", user.Register)
		userGroup.POST("/login", user.Login)

		userGroup.POST("/assign-role", user.AuthMiddleware(), user.AssignRole)
		userGroup.POST("/remove-role", user.AuthMiddleware(), user.RemoveRole)

		userGroup.GET("/me", user.AuthMiddleware(), user.GetMe)
		userGroup.PUT("/profile", user.AuthMiddleware(), user.UpdateProfile)
		userGroup.GET("/all", user.AuthMiddleware(), user.GetAllUsers)
		userGroup.GET("/search", user.AuthMiddleware(), user.SearchUsers)

		userGroup.GET("/:id/roles", user.AuthMiddleware(), user.ListRoles)
		userGroup.GET("/:id", user.AuthMiddleware(), user.GetProfile)
		userGroup.DELETE("/:id", user.AuthMiddleware(), user.DeleteUser)
	}
	bookingGroup := router.Group("/bookings")
	{
		bookingGroup.POST("/", booking.CreateBooking)
		bookingGroup.DELETE("/:id", booking.CancelBooking)
		bookingGroup.GET("/", booking.ListClientBookings)
	}
	router.GET("/available-slots", booking.ListAvailableSlots)

	router.Run(":8080")
}
