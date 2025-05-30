package application

import "github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/domain/entity"

type SalonUsecase interface {
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

	AssignProcedureToSpecialist(specialistID, procedureID string) error
	RemoveProcedureFromSpecialist(specialistID, procedureID string) error
}
