package domain

import "context"

type Payment struct {
	ID            string
	UserID        string
	Amount        int64
	Currency      string
	PaymentMethod string
	SalonID       string
	ServiceID     string
	StripeID      string
	Status        string
	CreatedAt     string
}

type PaymentRepository interface {
	Save(ctx context.Context, p *Payment) error
	ListByUser(ctx context.Context, userID string) ([]*Payment, error)
}
