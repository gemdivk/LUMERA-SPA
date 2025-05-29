package cache

import "github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/domain/entity"

type SalonCache interface {
	SetProcedures([]*entity.Procedure)
	GetProcedures() ([]*entity.Procedure, bool)
	SetSpecialists([]*entity.Specialist)
	GetSpecialists() ([]*entity.Specialist, bool)
}
