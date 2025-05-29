package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/gemdivk/LUMERA-SPA/notification-service/internal/domain"
	"github.com/stretchr/testify/assert"
)

// 🔧 Моки

type MockSender struct {
	SentTo     string
	SentSubj   string
	SentBody   string
	ShouldFail bool
}

func (m *MockSender) Send(to, subject, body string) error {
	if m.ShouldFail {
		return errors.New("smtp error")
	}
	m.SentTo = to
	m.SentSubj = subject
	m.SentBody = body
	return nil
}

type MockRepo struct {
	Logs       []domain.EmailLog
	LogCalled  bool
	ShouldFail bool
}

func (m *MockRepo) LogEmail(email, subject string) error {
	m.LogCalled = true
	if m.ShouldFail {
		return errors.New("db error")
	}
	m.Logs = append(m.Logs, domain.EmailLog{
		ID:      1,
		Email:   email,
		Subject: subject,
		SentAt:  time.Now(),
	})
	return nil
}

func (m *MockRepo) FetchAll() ([]domain.EmailLog, error) {
	return m.Logs, nil
}

// ✅ Тесты

func TestSendVerificationEmail_Success(t *testing.T) {
	sender := &MockSender{}
	repo := &MockRepo{}
	uc := NewEmailUsecase(sender, repo)

	err := uc.SendVerificationEmail("test@example.com", "token123")
	assert.NoError(t, err)
	assert.True(t, repo.LogCalled)
	assert.Equal(t, "test@example.com", sender.SentTo)
	assert.Contains(t, sender.SentBody, "token123")
	assert.Equal(t, "Email Verification", sender.SentSubj)
}

func TestSendVerificationEmail_SendFails(t *testing.T) {
	sender := &MockSender{ShouldFail: true}
	repo := &MockRepo{}
	uc := NewEmailUsecase(sender, repo)

	err := uc.SendVerificationEmail("fail@example.com", "tok")
	assert.Error(t, err)
	assert.False(t, repo.LogCalled) // лог не должен вызываться, если отправка упала
}

func TestSendVerificationEmail_LogFails(t *testing.T) {
	sender := &MockSender{}
	repo := &MockRepo{ShouldFail: true}
	uc := NewEmailUsecase(sender, repo)

	err := uc.SendVerificationEmail("test@example.com", "token")
	assert.Error(t, err)
	assert.True(t, repo.LogCalled)
}

func TestGetLogs(t *testing.T) {
	sender := &MockSender{}
	repo := &MockRepo{
		Logs: []domain.EmailLog{
			{ID: 1, Email: "a@a.com", Subject: "test", SentAt: time.Now()},
		},
	}
	uc := NewEmailUsecase(sender, repo)

	logs, err := uc.GetLogs()
	assert.NoError(t, err)
	assert.Len(t, logs, 1)
	assert.Equal(t, "a@a.com", logs[0].Email)
}
