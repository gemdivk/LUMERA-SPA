package booking

import (
	"net/http"

	pb "github.com/gemdivk/LUMERA-SPA/booking-service/proto"
	"github.com/gin-gonic/gin"
)

func CreateBooking(c *gin.Context) {
	var req pb.CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	resp, err := BookingClient.CreateBooking(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func CancelBooking(c *gin.Context) {
	id := c.Param("id")
	_, err := BookingClient.CancelBooking(c, &pb.CancelBookingRequest{BookingId: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func RescheduleBooking(c *gin.Context) {
	id := c.Param("id")
	var req pb.RescheduleBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	req.BookingId = id
	resp, err := BookingClient.RescheduleBooking(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ListClientBookings(c *gin.Context) {
	clientID := c.Query("client_id")
	resp, err := BookingClient.ListBookings(c, &pb.ListBookingsRequest{ClientId: clientID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Bookings)
}

func ListAllBookings(c *gin.Context) {
	resp, err := BookingClient.GetAllBookings(c, &pb.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Bookings)
}
