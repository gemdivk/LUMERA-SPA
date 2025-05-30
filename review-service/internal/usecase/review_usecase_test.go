package usecase

import (
	"github.com/gemdivk/LUMERA-SPA/review-service/test/mocks"
	"testing"
	"time"

	"github.com/gemdivk/LUMERA-SPA/review-service/internal/domain"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestReviewInteractor_CreateReview_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockReviewRepo(ctrl)
	interactor := NewReviewInteractor(mockRepo)

	review := &domain.Review{
		SalonID: "salon-1",
		UserID:  "user-1",
		Content: "Very good!",
		Rating:  5,
	}

	mockRepo.EXPECT().Create(gomock.Any()).Return(review, nil)

	result, err := interactor.CreateReview(review)

	assert.NoError(t, err)
	assert.Equal(t, review, result)
}

func TestReviewInteractor_CreateReview_ValidationError(t *testing.T) {
	interactor := NewReviewInteractor(nil)

	tests := []struct {
		name   string
		input  *domain.Review
		expect string
	}{
		{
			name:   "invalid rating",
			input:  &domain.Review{Rating: 0, Content: "test"},
			expect: "rating must be between 1 and 5",
		},
		{
			name:   "empty content",
			input:  &domain.Review{Rating: 3, Content: ""},
			expect: "review content cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := interactor.CreateReview(tt.input)
			assert.Nil(t, res)
			assert.EqualError(t, err, tt.expect)
		})
	}
}

func TestReviewInteractor_GetReview_CacheHit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockReviewRepo(ctrl)
	interactor := NewReviewInteractor(mockRepo)

	review := &domain.Review{
		ID:      "review-1",
		Content: "Cached review",
		Rating:  4,
	}

	// Установим вручную в кэш
	interactor.cache.Set(review.ID, review, time.Minute)

	result, err := interactor.GetReview("review-1")

	assert.NoError(t, err)
	assert.Equal(t, review, result)
}

func TestReviewInteractor_GetReview_DBHit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockReviewRepo(ctrl)
	interactor := NewReviewInteractor(mockRepo)

	expected := &domain.Review{
		ID:      "review-2",
		Content: "From DB",
		Rating:  3,
	}

	mockRepo.EXPECT().Get("review-2").Return(expected, nil)

	result, err := interactor.GetReview("review-2")

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestReviewInteractor_UpdateReview_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockReviewRepo(ctrl)
	interactor := NewReviewInteractor(mockRepo)

	review := &domain.Review{
		ID:      "id-123",
		Content: "Updated",
		Rating:  4,
	}

	mockRepo.EXPECT().Update(review).Return(review, nil)

	updated, err := interactor.UpdateReview(review)

	assert.NoError(t, err)
	assert.Equal(t, review, updated)
}

func TestReviewInteractor_DeleteReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockReviewRepo(ctrl)
	interactor := NewReviewInteractor(mockRepo)

	mockRepo.EXPECT().Delete("id-999").Return(nil)

	err := interactor.DeleteReview("id-999")
	assert.NoError(t, err)
}
