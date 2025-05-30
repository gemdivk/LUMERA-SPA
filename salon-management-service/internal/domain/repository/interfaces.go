package repository

import "github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/domain/entity"

type SalonRepository interface {
	AddSalon(*entity.Salon) (*entity.Salon, error)
	UpdateSalon(*entity.Salon) error
	DeleteSalon(id string) error
	GetAllSalons() ([]*entity.Salon, error)

	AddProcedure(*entity.Procedure) (*entity.Procedure, error)
	UpdateProcedure(*entity.Procedure) error
	DeleteProcedure(id string) error
	GetAllProcedures() ([]*entity.Procedure, error)

	AddSpecialist(*entity.Specialist) (*entity.Specialist, error)
	UpdateSpecialist(*entity.Specialist) error
	DeleteSpecialist(id string) error
	GetAllSpecialists() ([]*entity.Specialist, error)

	GetScheduleOverride(procedureID string, date string) (*entity.ProcedureScheduleOverride, error)
	GetWeeklySchedule(procedureID string, weekday int32) (*entity.WeeklyProcedureSchedule, error)
	AssignProcedureToSpecialist(specialistID, procedureID string) error
	RemoveProcedureFromSpecialist(specialistID, procedureID string) error
}
