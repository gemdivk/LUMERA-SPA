package usecase

import (
	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain"
	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain/application"
	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain/repository"
)

type BookingInteractor struct {
	repo repository.BookingRepository
}

func NewBookingUsecase(repo repository.BookingRepository) application.BookingUsecase {
	return &BookingInteractor{repo}
}

func (u *BookingInteractor) Create(b *domain.Booking) (*domain.Booking, error) {
	return u.repo.Create(b)
}

func (u *BookingInteractor) Cancel(id string) error {
	return u.repo.Cancel(id)
}

func (u *BookingInteractor) Reschedule(id, date, time string) (*domain.Booking, error) {
	return u.repo.Reschedule(id, date, time)
}

func (u *BookingInteractor) ListByClient(clientID string) ([]*domain.Booking, error) {
	return u.repo.ListByClient(clientID)
}
