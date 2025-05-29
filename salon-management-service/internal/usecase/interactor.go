package usecase

import (
	"github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/domain/application"
	"github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/domain/cache"
	"github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/domain/entity"
	"github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/domain/repository"
)

type SalonInteractor struct {
	repo  repository.SalonRepository
	cache cache.SalonCache
}

func NewSalonUsecase(repo repository.SalonRepository, cache cache.SalonCache) application.SalonUsecase {
	return &SalonInteractor{repo, cache}
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

func (s *SalonInteractor) GetAllProcedures() ([]*entity.Procedure, error) {
	if cached, ok := s.cache.GetProcedures(); ok {
		return cached, nil
	}
	data, err := s.repo.GetAllProcedures()
	if err == nil {
		s.cache.SetProcedures(data)
	}
	return data, err
}

func (s *SalonInteractor) GetAllSpecialists() ([]*entity.Specialist, error) {
	if cached, ok := s.cache.GetSpecialists(); ok {
		return cached, nil
	}
	data, err := s.repo.GetAllSpecialists()
	if err == nil {
		s.cache.SetSpecialists(data)
	}
	return data, err
}

func (s *SalonInteractor) AssignProcedureToSpecialist(specialistID, procedureID string) error {
	return s.repo.AssignProcedureToSpecialist(specialistID, procedureID)
}
