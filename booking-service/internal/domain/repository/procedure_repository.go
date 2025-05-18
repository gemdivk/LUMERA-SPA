package repository

import "github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain"

type ProcedureRepository interface {
	Create(p *domain.Procedure) (*domain.Procedure, error)
	GetByID(id string) (*domain.Procedure, error)
}
