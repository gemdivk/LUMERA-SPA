package main

import (
	"github.com/gemdivk/LUMERA-SPA/api-gateway/review"
	"github.com/gin-gonic/gin"
)

func main() {
	review.InitGRPCClient()

	router := gin.Default()

	rg := router.Group("/reviews")
	{
		rg.POST("/", review.CreateReview)
		rg.GET("/:id", review.GetReview)
		rg.PUT("/:id", review.UpdateReview)
		rg.DELETE("/:id", review.DeleteReview)
	}
	router.GET("/salons/:salon_id/reviews", review.ListReviews)

	router.Run(":8080")
}
