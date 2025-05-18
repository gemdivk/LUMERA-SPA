package booking

import (
	"net/http"
	"time"

	pb "github.com/gemdivk/LUMERA-SPA/booking-service/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateBooking(c *gin.Context) {
	var req struct {
		ClientID     string `json:"client_id"`
		SpecialistID string `json:"specialist_id"`
		ProcedureID  string `json:"procedure_id"`
		StartTime    string `json:"start_time"` // RFC3339 format
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	start, err := time.Parse(time.RFC3339, req.StartTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start time format"})
		return
	}

	resp, err := BookingClient.CreateBooking(c, &pb.CreateBookingRequest{
		ClientId:     req.ClientID,
		SpecialistId: req.SpecialistID,
		ProcedureId:  req.ProcedureID,
		StartTime:    timestamppb.New(start),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func CancelBooking(c *gin.Context) {
	id := c.Param("id")
	resp, err := BookingClient.CancelBooking(c, &pb.CancelBookingRequest{BookingId: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ListClientBookings(c *gin.Context) {
	clientID := c.Query("client_id")
	if clientID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "client_id is required"})
		return
	}
	resp, err := BookingClient.ListClientBookings(c, &pb.ClientBookingsRequest{ClientId: clientID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ListAvailableSlots(c *gin.Context) {
	procedureID := c.Query("procedure_id")
	date := c.Query("date")
	if procedureID == "" || date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "procedure_id and date required"})
		return
	}

	resp, err := BookingClient.ListAvailableSlots(c, &pb.ListAvailableSlotsRequest{
		ProcedureId: procedureID,
		Date:        date,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
