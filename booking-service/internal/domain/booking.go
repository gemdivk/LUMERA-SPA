package domain

import "time"

type Booking struct {
	ID           string
	ClientID     string
	SpecialistID string
	ProcedureID  string
	StartTime    time.Time
	EndTime      time.Time
	Status       string
	CreatedAt    time.Time
}
