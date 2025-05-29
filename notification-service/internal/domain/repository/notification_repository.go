package repository

import "github.com/gemdivk/LUMERA-SPA/notification-service/internal/domain"

type EmailLogRepo interface {
	LogEmail(email, subject string) error
	FetchAll() ([]domain.EmailLog, error)
}
