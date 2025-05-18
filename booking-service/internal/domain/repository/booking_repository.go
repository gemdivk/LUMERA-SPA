package repository

import "github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain"

type BookingRepository interface {
	Create(booking *domain.Booking) (*domain.Booking, error)
	Cancel(id string) error
	GetByClient(clientID string) ([]*domain.Booking, error)
	IsSlotAvailable(specialistID string, start, end string) (bool, error)
}
