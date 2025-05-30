package usecase

import (
	"errors"
	"log"
	"time"

	"github.com/gemdivk/LUMERA-SPA/review-service/internal/domain"
	"github.com/gemdivk/LUMERA-SPA/review-service/internal/domain/repository"
	"github.com/patrickmn/go-cache"
)

type ReviewInteractor struct {
	Repo  repository.ReviewRepo
	cache *cache.Cache
}

func NewReviewInteractor(repo repository.ReviewRepo) *ReviewInteractor {
	c := cache.New(10*time.Minute, 20*time.Minute)
	return &ReviewInteractor{
		Repo:  repo,
		cache: c,
	}
}

func (u *ReviewInteractor) CreateReview(review *domain.Review) (*domain.Review, error) {
	if review.Rating < 1 || review.Rating > 5 {
		return nil, errors.New("rating must be between 1 and 5")
	}
	if review.Content == "" {
		return nil, errors.New("review content cannot be empty")
	}

	created, err := u.Repo.Create(review)
	if err == nil {
		log.Printf("[CACHE] Set review id=%s", created.ID)
		u.cache.Set(created.ID, created, cache.DefaultExpiration)
	}
	return created, err
}

func (u *ReviewInteractor) GetReview(id string) (*domain.Review, error) {
	if cached, found := u.cache.Get(id); found {
		log.Printf("[CACHE] Hit for id=%s", id)
		return cached.(*domain.Review), nil
	}

	log.Printf("[CACHE] Miss for id=%s", id)
	review, err := u.Repo.Get(id)
	if err == nil {
		u.cache.Set(id, review, cache.DefaultExpiration)
		log.Printf("[CACHE] Cached after DB get for id=%s", id)
	}
	return review, err
}

func (u *ReviewInteractor) UpdateReview(review *domain.Review) (*domain.Review, error) {
	if review.Rating < 1 || review.Rating > 5 {
		return nil, errors.New("rating must be between 1 and 5")
	}
	if review.Content == "" {
		return nil, errors.New("review content cannot be empty")
	}

	updated, err := u.Repo.Update(review)
	if err == nil {
		u.cache.Set(review.ID, updated, cache.DefaultExpiration)
		log.Printf("[CACHE] Updated cache for id=%s", review.ID)
	}
	return updated, err
}

func (u *ReviewInteractor) DeleteReview(id string) error {
	err := u.Repo.Delete(id)
	if err == nil {
		u.cache.Delete(id)
		log.Printf("[CACHE] Deleted from cache id=%s", id)
	}
	return err
}

func (u *ReviewInteractor) ListBySalon(salonID string) ([]*domain.Review, error) {
	// Пока кэшировать весь список по salonID не будем — потенциально большой объём
	return u.Repo.ListBySalon(salonID)
}
