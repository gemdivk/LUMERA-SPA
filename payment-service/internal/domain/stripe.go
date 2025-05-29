package domain

import "time"

type PaymentIntentInfo struct {
	ID        string
	Status    string
	CreatedAt time.Time
}

type StripeClient interface {
	CreateIntent(amount int64, currency, method string) (*PaymentIntentInfo, error)
}
