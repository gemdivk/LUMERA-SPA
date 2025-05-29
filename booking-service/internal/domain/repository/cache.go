package repository

import "github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain"

type BookingCache interface {
	GetClientBookings(clientID string) ([]*domain.Booking, bool)
	SetClientBookings(clientID string, bookings []*domain.Booking)

	GetAllBookings() ([]*domain.Booking, bool)
	SetAllBookings(bookings []*domain.Booking)
}
