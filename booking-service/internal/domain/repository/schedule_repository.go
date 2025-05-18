package repository

import "github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain"

type ScheduleRepository interface {
	CreateTemplate(t *domain.ScheduleTemplate) error
	GetTemplateByWeekday(specialistID string, weekday int) (*domain.ScheduleTemplate, error)

	OverrideDailySchedule(s *domain.DailySchedule) error
	GetDailySchedule(specialistID string, date string) (*domain.DailySchedule, error)

	GetAssignedProcedures(specialistID string) ([]*domain.Procedure, error)
	GetAvailableSlots(procedureID, date string) ([]domain.TimeSlot, error)
	AssignProcedure(specialistID, procedureID string) error
}
