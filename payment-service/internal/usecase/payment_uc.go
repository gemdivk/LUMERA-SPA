package usecase

import (
	"context"
	"encoding/json"
	"github.com/gemdivk/LUMERA-SPA/payment-service/internal/domain"
	"time"
)

type PaymentUsecase struct {
	Repo   domain.PaymentRepository
	Tx     domain.TxRunner
	Stripe domain.StripeClient
	NATS   domain.NATSClient
}

func New(
	repo domain.PaymentRepository,
	txRunner domain.TxRunner,
	stripeClient domain.StripeClient,
	natsClient domain.NATSClient,
) *PaymentUsecase {
	return &PaymentUsecase{
		Repo:   repo,
		Tx:     txRunner,
		Stripe: stripeClient,
		NATS:   natsClient,
	}
}

func (uc *PaymentUsecase) Create(ctx context.Context, userID string, amount int64, currency string, method string, salonID string, serviceID string) (*domain.Payment, error) {

	info, err := uc.Stripe.CreateIntent(amount, currency, method)
	if err != nil {
		return nil, err
	}

	p := &domain.Payment{
		ID:            info.ID,
		UserID:        userID,
		Amount:        amount,
		Currency:      currency,
		SalonID:       salonID,
		ServiceID:     serviceID,
		PaymentMethod: method,
		StripeID:      info.ID,
		Status:        string(info.Status),
		CreatedAt:     info.CreatedAt.Format(time.RFC3339),
	}

	if err1 := uc.Tx.RunTx(ctx, func(repo domain.PaymentRepository) error {
		if err := repo.Save(ctx, p); err != nil {
			return err
		}
		b, _ := json.Marshal(p)
		return uc.NATS.Publish("payment.created", b)
	}); err1 != nil {
		return nil, err1
	}

	return p, nil
}

func (uc *PaymentUsecase) List(ctx context.Context, userID string) ([]*domain.Payment, error) {
	return uc.Repo.ListByUser(ctx, userID)
}
