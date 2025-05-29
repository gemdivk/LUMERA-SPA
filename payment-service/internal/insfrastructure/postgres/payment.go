package repository

import (
	"context"
	"github.com/gemdivk/LUMERA-SPA/payment-service/internal/domain"

	"gorm.io/gorm"
)

type PgRepo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *PgRepo {
	return &PgRepo{DB: db}
}

func (r *PgRepo) WithTx(tx *gorm.DB) *PgRepo {
	return &PgRepo{DB: tx}
}

func (r *PgRepo) RunTx(ctx context.Context, fn func(repo domain.PaymentRepository) error) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txRepo := r.WithTx(tx)
		return fn(txRepo)
	})
}

func (r *PgRepo) Save(ctx context.Context, p *domain.Payment) error {
	return r.DB.WithContext(ctx).Create(p).Error
}

func (r *PgRepo) ListByUser(ctx context.Context, userID string) ([]*domain.Payment, error) {
	var payments []*domain.Payment
	err := r.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&payments).Error
	return payments, err
}

func (r *PgRepo) GetByID(ctx context.Context, id string) (*domain.Payment, error) {
	var p domain.Payment
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}
