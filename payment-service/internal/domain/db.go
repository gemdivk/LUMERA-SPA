package domain

import (
	"context"
)

type TxRunner interface {
	RunTx(ctx context.Context, fn func(repo PaymentRepository) error) error
}
