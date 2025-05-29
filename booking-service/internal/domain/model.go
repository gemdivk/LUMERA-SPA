package domain

type Booking struct {
	ID           string
	ClientID     string
	SalonID      string
	ProcedureID  string
	SpecialistID string
	Date         string
	StartTime    string
	Status       string
}
