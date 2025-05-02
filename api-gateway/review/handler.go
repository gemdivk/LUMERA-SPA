package review

import (
	"net/http"

	pb "github.com/gemdivk/LUMERA-SPA/review-service/proto"
	"github.com/gin-gonic/gin"
)

func CreateReview(c *gin.Context) {
	var req pb.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный JSON"})
		return
	}
	resp, err := ReviewClient.CreateReview(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func GetReview(c *gin.Context) {
	id := c.Param("id")
	resp, err := ReviewClient.GetReview(c, &pb.GetReviewRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ListReviews(c *gin.Context) {
	id := c.Param("salon_id")
	resp, err := ReviewClient.ListReviews(c, &pb.ListReviewsRequest{SalonId: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func UpdateReview(c *gin.Context) {
	id := c.Param("id")
	var req pb.UpdateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный JSON"})
		return
	}
	req.Id = id
	resp, err := ReviewClient.UpdateReview(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	_, err := ReviewClient.DeleteReview(c, &pb.DeleteReviewRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
