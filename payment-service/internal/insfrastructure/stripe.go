package infrastructure

import (
	"github.com/gemdivk/LUMERA-SPA/payment-service/internal/domain"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/paymentintent"
	"log"
	"os"
	"time"
)

type StripeAdapter struct{}

func InitStripeClient() domain.StripeClient {

	key := os.Getenv("STRIPE_SECRET_KEY")
	if key == "" {
		log.Fatal("STRIPE_SECRET_KEY is not set")
	}
	stripe.Key = key
	return &StripeAdapter{}
}

func (s *StripeAdapter) CreateIntent(amount int64, currency, method string) (*domain.PaymentIntentInfo, error) {
	params := &stripe.PaymentIntentParams{
		Amount:        stripe.Int64(amount),
		Currency:      stripe.String(currency),
		Confirm:       stripe.Bool(true),
		PaymentMethod: stripe.String(method),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled:        stripe.Bool(true),
			AllowRedirects: stripe.String("never"),
		},
	}
	pi, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	return &domain.PaymentIntentInfo{
		ID:        pi.ID,
		Status:    string(pi.Status),
		CreatedAt: time.Unix(pi.Created, 0),
	}, nil
}
