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

func (s *SalonInteractor) UpdateSalon(salon *entity.Salon) error {
	return s.repo.UpdateSalon(salon)
}

func (s *SalonInteractor) DeleteSalon(id string) error {
	return s.repo.DeleteSalon(id)
}

func (s *SalonInteractor) GetAllSalons() ([]*entity.Salon, error) {
	return s.repo.GetAllSalons()
}

func (s *SalonInteractor) UpdateProcedure(p *entity.Procedure) error {
	return s.repo.UpdateProcedure(p)
}

func (s *SalonInteractor) DeleteProcedure(id string) error {
	return s.repo.DeleteProcedure(id)
}

func (s *SalonInteractor) UpdateSpecialist(sp *entity.Specialist) error {
	return s.repo.UpdateSpecialist(sp)
}

func (s *SalonInteractor) DeleteSpecialist(id string) error {
	return s.repo.DeleteSpecialist(id)
}

func (s *SalonInteractor) RemoveProcedureFromSpecialist(specialistID, procedureID string) error {
	return s.repo.RemoveProcedureFromSpecialist(specialistID, procedureID)
}
