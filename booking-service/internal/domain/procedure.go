package domain

import "time"

type Procedure struct {
	ID              string
	Name            string
	DurationMinutes int32
	CreatedAt       time.Time
}
