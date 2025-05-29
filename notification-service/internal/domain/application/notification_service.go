package application

import "github.com/gemdivk/LUMERA-SPA/notification-service/internal/domain"

type EmailUsecase interface {
	SendVerificationEmail(email, token string) error
	GetLogs() ([]domain.EmailLog, error)
}

type EmailSender interface {
	Send(to, subject, body string) error
}
