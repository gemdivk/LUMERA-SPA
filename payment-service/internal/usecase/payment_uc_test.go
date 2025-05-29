package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/gemdivk/LUMERA-SPA/payment-service/internal/domain"
	"github.com/gemdivk/LUMERA-SPA/payment-service/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- Моки ---

// MockRepo реализует domain.PaymentRepository
type MockRepo struct{ mock.Mock }

func (m *MockRepo) Save(ctx context.Context, p *domain.Payment) error {
	return m.Called(ctx, p).Error(0)
}

func (m *MockRepo) ListByUser(ctx context.Context, userID string) ([]*domain.Payment, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]*domain.Payment), args.Error(1)
}

func (m *MockRepo) GetByID(ctx context.Context, id string) (*domain.Payment, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Payment), args.Error(1)
}

// FakeTxRunner реализует domain.TxRunner — сразу выполняет fn(mockRepo)
type FakeTxRunner struct {
	Repo domain.PaymentRepository
}

func (f *FakeTxRunner) RunTx(ctx context.Context, fn func(repo domain.PaymentRepository) error) error {
	return fn(f.Repo)
}

// MockStripe реализует domain.StripeClient
type MockStripe struct{ mock.Mock }

func (m *MockStripe) CreateIntent(amount int64, currency, method string) (*domain.PaymentIntentInfo, error) {
	args := m.Called(amount, currency, method)
	return args.Get(0).(*domain.PaymentIntentInfo), args.Error(1)
}

// MockNATS реализует domain.NATSClient
type MockNATS struct{ mock.Mock }

func (m *MockNATS) Publish(subject string, data []byte) error {
	return m.Called(subject, data).Error(0)
}

// --- Тесты ---

func TestCreatePayment_Success(t *testing.T) {
	// Подготавливаем моки
	mockRepo := new(MockRepo)
	txRunner := &FakeTxRunner{Repo: mockRepo}
	mockStripe := new(MockStripe)
	mockNATS := new(MockNATS)

	uc := usecase.New(mockRepo, txRunner, mockStripe, mockNATS)

	// Входные данные
	amount := int64(1200)
	currency := "usd"
	method := "pm_card_visa"
	userID := "u1"
	salonID := "s1"
	serviceID := "svc1"

	// Ответ от Stripe
	now := time.Now()
	intentInfo := &domain.PaymentIntentInfo{
		ID:        "pi_123",
		Status:    "succeeded",
		CreatedAt: now,
	}
	mockStripe.
		On("CreateIntent", amount, currency, method).
		Return(intentInfo, nil)

	// Ожидаем вызов Save
	mockRepo.
		On("Save", mock.Anything, mock.MatchedBy(func(p *domain.Payment) bool {
			return p.UserID == userID &&
				p.Amount == amount &&
				p.ID == intentInfo.ID &&
				p.Status == intentInfo.Status
		})).
		Return(nil)

	// Ожидаем вызов Publish
	mockNATS.
		On("Publish", "payment.created", mock.Anything).
		Return(nil)

	// Запускаем
	out, err := uc.Create(context.Background(), userID, amount, currency, method, salonID, serviceID)
	assert.NoError(t, err)
	assert.Equal(t, intentInfo.ID, out.ID)
	assert.Equal(t, intentInfo.Status, out.Status)

	// Проверяем, что все ожидания выполнены
	mockStripe.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockNATS.AssertExpectations(t)
}

func TestCreatePayment_RepoError(t *testing.T) {
	mockRepo := new(MockRepo)
	txRunner := &FakeTxRunner{Repo: mockRepo}
	mockStripe := new(MockStripe)
	mockNATS := new(MockNATS)

	uc := usecase.New(mockRepo, txRunner, mockStripe, mockNATS)

	// Stripe OK
	intentInfo := &domain.PaymentIntentInfo{
		ID:        "pi_err",
		Status:    "succeeded",
		CreatedAt: time.Now(),
	}
	mockStripe.
		On("CreateIntent", mock.Anything, mock.Anything, mock.Anything).
		Return(intentInfo, nil)

	// Репо возвращает ошибку при Save
	mockRepo.
		On("Save", mock.Anything, mock.Anything).
		Return(errors.New("db-error"))

	// Запускаем
	_, err := uc.Create(context.Background(), "u1", 500, "usd", "pm", "s", "svc")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "db-error")

	// Проверяем ожидания
	mockStripe.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}
