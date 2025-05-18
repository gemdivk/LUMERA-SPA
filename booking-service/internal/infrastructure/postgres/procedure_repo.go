package postgres

import (
	"database/sql"
	"time"

	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain"
	"github.com/google/uuid"
)

type ProcedureRepo struct {
	DB *sql.DB
}

func NewProcedureRepo(db *sql.DB) *ProcedureRepo {
	return &ProcedureRepo{DB: db}
}

func (r *ProcedureRepo) Create(p *domain.Procedure) (*domain.Procedure, error) {
	p.ID = uuid.New().String()
	p.CreatedAt = time.Now()
	_, err := r.DB.Exec(`INSERT INTO procedures (id, name, duration_minutes, created_at) VALUES ($1, $2, $3, $4)`,
		p.ID, p.Name, p.DurationMinutes, p.CreatedAt)
	return p, err
}

func (r *ProcedureRepo) GetByID(id string) (*domain.Procedure, error) {
	row := r.DB.QueryRow(`SELECT id, name, duration_minutes, created_at FROM procedures WHERE id = $1`, id)
	p := &domain.Procedure{}
	err := row.Scan(&p.ID, &p.Name, &p.DurationMinutes, &p.CreatedAt)
	return p, err
}
