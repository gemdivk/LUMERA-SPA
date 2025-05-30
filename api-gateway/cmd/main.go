package main

import (
	"github.com/gemdivk/LUMERA-SPA/api-gateway/booking"
	"github.com/gemdivk/LUMERA-SPA/api-gateway/notification"
	"github.com/gemdivk/LUMERA-SPA/api-gateway/review"
	"github.com/gemdivk/LUMERA-SPA/api-gateway/salon"
	"github.com/gemdivk/LUMERA-SPA/api-gateway/user"
	"github.com/gin-gonic/gin"
)

func main() {
	review.InitGRPCClient()
	user.InitGRPCClient()
	notification.InitGRPCClient()
	booking.InitGRPCClient()
	salon.InitGRPCClient()

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
		userGroup.POST("/logout", user.AuthMiddleware(), user.Logout)

		userGroup.POST("/assign-role", user.AuthMiddleware(), user.AssignRole)
		userGroup.POST("/remove-role", user.AuthMiddleware(), user.RemoveRole)

		userGroup.GET("/me", user.AuthMiddleware(), user.GetMe)
		userGroup.PUT("/profile", user.AuthMiddleware(), user.UpdateProfile)
		userGroup.GET("/all", user.AuthMiddleware(), user.GetAllUsers)
		userGroup.GET("/search", user.AuthMiddleware(), user.SearchUsers)

		userGroup.GET("/:id/roles", user.AuthMiddleware(), user.ListRoles)
		userGroup.GET("/:id", user.AuthMiddleware(), user.GetProfile)
		userGroup.DELETE("/:id", user.AuthMiddleware(), user.DeleteUser)

		router.GET("/verify", user.VerifyEmail)
	}

	bookingGroup := router.Group("/bookings")
	{
		bookingGroup.POST("/", booking.CreateBooking)
		bookingGroup.PUT("/:id/reschedule", booking.RescheduleBooking)
		bookingGroup.DELETE("/:id", booking.CancelBooking)

		bookingGroup.GET("/", booking.ListAllBookings)
		bookingGroup.GET("/by-client", booking.ListClientBookings)
	}

	salonGroup := router.Group("/salon")
	{
		salonGroup.POST("/", salon.AddSalon)
		salonGroup.PUT("/:id", salon.UpdateSalon)
		salonGroup.DELETE("/:id", salon.DeleteSalon)
		salonGroup.GET("/", salon.GetAllSalons)

		salonGroup.POST("/procedures", salon.AddProcedure)
		salonGroup.PUT("/procedures/:id", salon.UpdateProcedure)
		salonGroup.DELETE("/procedures/:id", salon.DeleteProcedure)
		salonGroup.GET("/procedures", salon.GetAllProcedures)
		salonGroup.GET("/procedures-by-time", salon.GetAllProceduresByTime)

		salonGroup.POST("/specialists", salon.AddSpecialist)
		salonGroup.PUT("/specialists/:id", salon.UpdateSpecialist)
		salonGroup.DELETE("/specialists/:id", salon.DeleteSpecialist)
		salonGroup.GET("/specialists", salon.GetAllSpecialists)

		salonGroup.POST("/assign-procedure", salon.AssignProcedureToSpecialist)
		salonGroup.POST("/unassign-procedure", salon.UnassignProcedureFromSpecialist)
	}
	notifications := router.Group("/notifications")
	{
		notifications.GET("/logs", user.AuthMiddleware(), notification.GetLogs)
	}

	router.Run(":8080")
}
