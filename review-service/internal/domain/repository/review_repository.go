package repository

import "github.com/gemdivk/LUMERA-SPA/review-service/internal/domain"

type ReviewRepo interface {
	Create(review *domain.Review) (*domain.Review, error)
	Get(id string) (*domain.Review, error)
	Update(review *domain.Review) (*domain.Review, error)
	Delete(id string) error
	ListBySalon(salonID string) ([]*domain.Review, error)
}
