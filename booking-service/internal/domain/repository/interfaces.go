package repository

import "github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain"

type BookingRepository interface {
	Create(*domain.Booking) (*domain.Booking, error)
	Cancel(string) error
	Reschedule(string, string, string) (*domain.Booking, error)
	ListByClient(string) ([]*domain.Booking, error)
	GetAll() ([]*domain.Booking, error)
}
