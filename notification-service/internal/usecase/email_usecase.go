package usecase

import (
	"github.com/gemdivk/LUMERA-SPA/notification-service/internal/domain"
	"github.com/gemdivk/LUMERA-SPA/notification-service/internal/domain/application"
	"github.com/gemdivk/LUMERA-SPA/notification-service/internal/domain/repository"
)

type EmailUsecase struct {
	Sender application.EmailSender
	Repo   repository.EmailLogRepo
}

func NewEmailUsecase(sender application.EmailSender, repo repository.EmailLogRepo) *EmailUsecase {
	return &EmailUsecase{Sender: sender, Repo: repo}
}

func (u *EmailUsecase) SendVerificationEmail(email, token string) error {
	subject := "Email Verification"
	body := "Please verify your email by clicking this link: http://localhost:8080/verify?token=" + token
	if err := u.Sender.Send(email, subject, body); err != nil {
		return err
	}
	return u.Repo.LogEmail(email, subject)
}

func (u *EmailUsecase) GetLogs() ([]domain.EmailLog, error) {
	return u.Repo.FetchAll()
}
