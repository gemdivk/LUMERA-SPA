package usecase

import (
	"github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/domain/application"
	"github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/domain/entity"
	"github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/domain/repository"
)

type SalonInteractor struct {
	repo repository.SalonRepository
}

func NewSalonUsecase(repo repository.SalonRepository) application.SalonUsecase {
	return &SalonInteractor{repo}
}

func (s *SalonInteractor) AddSalon(salon *entity.Salon) (*entity.Salon, error) {
	return s.repo.AddSalon(salon)
}

func (s *SalonInteractor) AddProcedure(proc *entity.Procedure) (*entity.Procedure, error) {
	return s.repo.AddProcedure(proc)
}

func (s *SalonInteractor) AddSpecialist(sp *entity.Specialist) (*entity.Specialist, error) {
	return s.repo.AddSpecialist(sp)
}
