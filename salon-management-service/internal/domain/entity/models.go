package entity

type Salon struct {
	ID       string
	Name     string
	Location string
}

type Procedure struct {
	ID          string
	SalonID     string
	Name        string
	Duration    int32
	Description string
}

type Specialist struct {
	ID      string
	SalonID string
	Name    string
	Bio     string
}

type WeeklyProcedureSchedule struct {
	DayOfWeek string
	StartTime string
	EndTime   string
}

type ProcedureScheduleOverride struct {
	Date      string
	StartTime string
	EndTime   string
}
