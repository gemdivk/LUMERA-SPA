package postgres

import (
	"database/sql"

	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain"
)

type ScheduleRepo struct {
	DB *sql.DB
}

func NewScheduleRepo(db *sql.DB) *ScheduleRepo {
	return &ScheduleRepo{DB: db}
}

func (r *ScheduleRepo) CreateTemplate(t *domain.ScheduleTemplate) error {
	return nil
}

func (r *ScheduleRepo) GetTemplateByWeekday(specialistID string, weekday int) (*domain.ScheduleTemplate, error) {
	return nil, nil
}

func (r *ScheduleRepo) OverrideDailySchedule(s *domain.DailySchedule) error {
	return nil
}

func (r *ScheduleRepo) GetDailySchedule(specialistID string, date string) (*domain.DailySchedule, error) {
	return nil, nil
}

func (r *ScheduleRepo) GetAssignedProcedures(specialistID string) ([]*domain.Procedure, error) {
	return nil, nil
}

func (r *ScheduleRepo) GetAvailableSlots(procedureID, date string) ([]domain.TimeSlot, error) {
	return nil, nil
}

func (r *ScheduleRepo) AssignProcedure(specialistID, procedureID string) error {
	_, err := r.DB.Exec(`INSERT INTO specialist_procedures (id, specialist_id, procedure_id) VALUES (gen_random_uuid(), $1, $2)`, specialistID, procedureID)
	return err
}
