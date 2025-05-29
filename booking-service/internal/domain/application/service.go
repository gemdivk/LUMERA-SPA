package application

import "github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain"

type BookingUsecase interface {
	Create(*domain.Booking) (*domain.Booking, error)
	Cancel(string) error
	Reschedule(string, string, string) (*domain.Booking, error)
	ListByClient(string) ([]*domain.Booking, error)
}
