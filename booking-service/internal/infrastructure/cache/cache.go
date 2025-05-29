package cache

import (
	"sync"

	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain"
)

type BookingCache struct {
	data map[string]*domain.Booking
	mu   sync.RWMutex
}

func NewBookingCache() *BookingCache {
	return &BookingCache{
		data: make(map[string]*domain.Booking),
	}
}

func (c *BookingCache) Set(b *domain.Booking) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[b.ID] = b
}

func (c *BookingCache) GetAll() []*domain.Booking {
	c.mu.RLock()
	defer c.mu.RUnlock()
	bookings := make([]*domain.Booking, 0, len(c.data))
	for _, b := range c.data {
		bookings = append(bookings, b)
	}
	return bookings
}

func (c *BookingCache) GetByClient(clientID string) []*domain.Booking {
	c.mu.RLock()
	defer c.mu.RUnlock()
	var result []*domain.Booking
	for _, b := range c.data {
		if b.ClientID == clientID {
			result = append(result, b)
		}
	}
	return result
}

func (c *BookingCache) Delete(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, id)
}
