package usecase

import (
	"log"

	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain"
	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain/application"
	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain/repository"
	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/infrastructure/cache"
)

type BookingInteractor struct {
	repo  repository.BookingRepository
	cache *cache.BookingCache
}

func NewBookingUsecase(repo repository.BookingRepository) application.BookingUsecase {
	c := cache.NewBookingCache()

	bookings, err := repo.GetAll()
	if err != nil {
		log.Printf("failed to load cache from DB: %v", err)
	} else {
		for _, b := range bookings {
			c.Set(b)
		}
	}

	return &BookingInteractor{
		repo:  repo,
		cache: c,
	}
}

func (u *BookingInteractor) Create(b *domain.Booking) (*domain.Booking, error) {
	created, err := u.repo.Create(b)
	if err == nil {
		u.cache.Set(created)
	}
	return created, err
}

func (u *BookingInteractor) Cancel(id string) error {
	err := u.repo.Cancel(id)
	if err == nil {
		for _, b := range u.cache.GetAll() {
			if b.ID == id {
				b.Status = "cancelled"
			}
		}
	}
	return err
}

func (u *BookingInteractor) Reschedule(id, date, time string) (*domain.Booking, error) {
	updated, err := u.repo.Reschedule(id, date, time)
	if err == nil {
		u.cache.Set(updated)
	}
	return updated, err
}

func (u *BookingInteractor) ListByClient(clientID string) ([]*domain.Booking, error) {
	return u.cache.GetByClient(clientID), nil
}

func (u *BookingInteractor) GetAllBookings() ([]*domain.Booking, error) {
	return u.cache.GetAll(), nil
}
