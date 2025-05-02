package usecase

import (
	"errors"
	"github.com/gemdivk/LUMERA-SPA/review-service/internal/domain"
	"github.com/gemdivk/LUMERA-SPA/review-service/internal/domain/repository"
)

type ReviewInteractor struct {
	Repo repository.ReviewRepo
}

func NewReviewInteractor(repo repository.ReviewRepo) *ReviewInteractor {
	return &ReviewInteractor{Repo: repo}
}

func (u *ReviewInteractor) CreateReview(review *domain.Review) (*domain.Review, error) {
	if review.Rating < 1 || review.Rating > 5 {
		return nil, errors.New("rating must be between 1 and 5")
	}
	if review.Content == "" {
		return nil, errors.New("review content cannot be empty")
	}
	return u.Repo.Create(review)
}

func (u *ReviewInteractor) GetReview(id string) (*domain.Review, error) {
	return u.Repo.Get(id)
}

func (u *ReviewInteractor) UpdateReview(review *domain.Review) (*domain.Review, error) {
	if review.Rating < 1 || review.Rating > 5 {
		return nil, errors.New("rating must be between 1 and 5")
	}
	if review.Content == "" {
		return nil, errors.New("review content cannot be empty")
	}
	return u.Repo.Update(review)
}

func (u *ReviewInteractor) DeleteReview(id string) error {
	return u.Repo.Delete(id)
}

func (u *ReviewInteractor) ListBySalon(salonID string) ([]*domain.Review, error) {
	return u.Repo.ListBySalon(salonID)
}
