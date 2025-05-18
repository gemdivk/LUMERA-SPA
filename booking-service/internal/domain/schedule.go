package domain

import "time"

type ScheduleTemplate struct {
	ID           string
	SpecialistID string
	Weekday      int
	StartTime    string
	EndTime      string
	BreakMinutes int32
}

type DailySchedule struct {
	ID           string
	SpecialistID string
	Date         time.Time
	StartTime    string
	EndTime      string
	Override     bool
	Cancelled    bool
}

type TimeSlot struct {
	SpecialistID   string
	SpecialistName string
	StartTime      time.Time
	EndTime        time.Time
}
