package usecase

import (
	"testing"

	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain"
	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepository(ctrl)

	booking := &domain.Booking{
		ID:           "1",
		ClientID:     "client1",
		SalonID:      "salon1",
		ProcedureID:  "proc1",
		SpecialistID: "spec1",
		Date:         "2025-06-01",
		StartTime:    "10:00",
		Status:       "booked",
	}

	mockRepo.EXPECT().GetAll().Return([]*domain.Booking{}, nil)

	mockRepo.EXPECT().
		Create(gomock.AssignableToTypeOf(&domain.Booking{})).
		Return(booking, nil)

	uc := NewBookingUsecase(mockRepo)
	result, err := uc.Create(booking)

	assert.NoError(t, err)
	assert.Equal(t, booking, result)
}

func TestCancel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepository(ctrl)

	mockRepo.EXPECT().GetAll().Return([]*domain.Booking{
		{
			ID:           "1",
			ClientID:     "client1",
			SalonID:      "salon1",
			ProcedureID:  "proc1",
			SpecialistID: "spec1",
			Date:         "2025-06-01",
			StartTime:    "10:00",
			Status:       "booked",
		},
	}, nil)

	mockRepo.EXPECT().Cancel("1").Return(nil)

	uc := NewBookingUsecase(mockRepo)
	err := uc.Cancel("1")

	assert.NoError(t, err)
}

func TestReschedule(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepository(ctrl)

	newBooking := &domain.Booking{
		ID:           "1",
		ClientID:     "client1",
		SalonID:      "salon1",
		ProcedureID:  "proc1",
		SpecialistID: "spec1",
		Date:         "2025-06-02",
		StartTime:    "12:00",
		Status:       "rescheduled",
	}

	mockRepo.EXPECT().GetAll().Return([]*domain.Booking{}, nil)
	mockRepo.EXPECT().Reschedule("1", "2025-06-02", "12:00").Return(newBooking, nil)

	uc := NewBookingUsecase(mockRepo)
	result, err := uc.Reschedule("1", "2025-06-02", "12:00")

	assert.NoError(t, err)
	assert.Equal(t, newBooking, result)
}
