package application

import "github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/domain/entity"

type SalonUsecase interface {
	AddSalon(*entity.Salon) (*entity.Salon, error)
	AddProcedure(*entity.Procedure) (*entity.Procedure, error)
	AddSpecialist(*entity.Specialist) (*entity.Specialist, error)
}
