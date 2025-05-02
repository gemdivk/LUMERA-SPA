package application

import "github.com/gemdivk/LUMERA-SPA/review-service/internal/domain"

type ReviewUsecase interface {
	CreateReview(review *domain.Review) (*domain.Review, error)
	GetReview(id string) (*domain.Review, error)
	UpdateReview(review *domain.Review) (*domain.Review, error)
	DeleteReview(id string) error
	ListBySalon(salonID string) ([]*domain.Review, error)
}
