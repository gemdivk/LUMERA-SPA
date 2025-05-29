package entity

type Salon struct {
	ID       string
	Name     string
	Location string
}

type Procedure struct {
	ID       string
	SalonID  string
	Name     string
	Duration int32
}

type Specialist struct {
	ID      string
	SalonID string
	Name    string
}
