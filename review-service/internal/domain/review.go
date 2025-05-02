package domain

import "time"

type Review struct {
	ID        string
	SalonID   string
	UserID    string
	Content   string
	Rating    int32
	CreatedAt time.Time
}
