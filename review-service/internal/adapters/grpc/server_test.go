package grpc

import (
	"context"
	"testing"
	"time"

	"github.com/gemdivk/LUMERA-SPA/review-service/internal/domain"
	grpc "github.com/gemdivk/LUMERA-SPA/review-service/proto"
	"github.com/gemdivk/LUMERA-SPA/review-service/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestReviewServer_CreateReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockReviewUsecase(ctrl)
	server := NewReviewServer(mockUsecase)

	req := &grpc.CreateReviewRequest{
		SalonId: "salon-1",
		UserId:  "user-1",
		Content: "Great service",
		Rating:  5,
	}

	expected := &domain.Review{
		ID:        "rev-1",
		SalonID:   "salon-1",
		UserID:    "user-1",
		Content:   "Great service",
		Rating:    5,
		CreatedAt: time.Now(),
	}

	mockUsecase.
		EXPECT().
		CreateReview(gomock.Any()).
		Return(expected, nil)

	resp, err := server.CreateReview(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, expected.ID, resp.Id)
	assert.Equal(t, expected.SalonID, resp.SalonId)
	assert.Equal(t, expected.UserID, resp.UserId)
	assert.Equal(t, expected.Content, resp.Content)
	assert.Equal(t, expected.Rating, resp.Rating)
}
func TestReviewServer_GetReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockReviewUsecase(ctrl)
	server := NewReviewServer(mockUsecase)

	expected := &domain.Review{
		ID:        "rev-1",
		SalonID:   "salon-1",
		UserID:    "user-1",
		Content:   "Nice place",
		Rating:    4,
		CreatedAt: time.Now(),
	}

	mockUsecase.
		EXPECT().
		GetReview("rev-1").
		Return(expected, nil)

	req := &grpc.GetReviewRequest{Id: "rev-1"}
	resp, err := server.GetReview(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, expected.ID, resp.Id)
	assert.Equal(t, expected.Content, resp.Content)
}
func TestReviewServer_UpdateReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockReviewUsecase(ctrl)
	server := NewReviewServer(mockUsecase)

	input := &domain.Review{
		ID:      "rev-1",
		Content: "Updated",
		Rating:  4,
	}
	expected := &domain.Review{
		ID:        "rev-1",
		SalonID:   "salon-1",
		UserID:    "user-1",
		Content:   "Updated",
		Rating:    4,
		CreatedAt: time.Now(),
	}

	mockUsecase.
		EXPECT().
		UpdateReview(input).
		Return(expected, nil)

	req := &grpc.UpdateReviewRequest{
		Id:      "rev-1",
		Content: "Updated",
		Rating:  4,
	}
	resp, err := server.UpdateReview(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, expected.Content, resp.Content)
	assert.Equal(t, expected.Rating, resp.Rating)
}
func TestReviewServer_DeleteReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockReviewUsecase(ctrl)
	server := NewReviewServer(mockUsecase)

	mockUsecase.
		EXPECT().
		DeleteReview("rev-1").
		Return(nil)

	req := &grpc.DeleteReviewRequest{Id: "rev-1"}
	resp, err := server.DeleteReview(context.Background(), req)

	assert.NoError(t, err)
	assert.True(t, resp.Success)
}
func TestReviewServer_ListReviews(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockReviewUsecase(ctrl)
	server := NewReviewServer(mockUsecase)

	list := []*domain.Review{
		{ID: "rev-1", SalonID: "salon-1", UserID: "user-1", Content: "Great", Rating: 5, CreatedAt: time.Now()},
		{ID: "rev-2", SalonID: "salon-1", UserID: "user-2", Content: "Good", Rating: 4, CreatedAt: time.Now()},
	}

	mockUsecase.
		EXPECT().
		ListBySalon("salon-1").
		Return(list, nil)

	req := &grpc.ListReviewsRequest{SalonId: "salon-1"}
	resp, err := server.ListReviews(context.Background(), req)

	assert.NoError(t, err)
	assert.Len(t, resp.Reviews, 2)
	assert.Equal(t, "rev-1", resp.Reviews[0].Id)
	assert.Equal(t, "rev-2", resp.Reviews[1].Id)
}
