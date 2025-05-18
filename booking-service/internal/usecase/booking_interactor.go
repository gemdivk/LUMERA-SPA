package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain"
	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain/repository"
)

type BookingInteractor struct {
	BookingRepo   repository.BookingRepository
	ProcedureRepo repository.ProcedureRepository
	ScheduleRepo  repository.ScheduleRepository
}

func NewBookingInteractor(b repository.BookingRepository, p repository.ProcedureRepository, s repository.ScheduleRepository) *BookingInteractor {
	return &BookingInteractor{
		BookingRepo:   b,
		ProcedureRepo: p,
		ScheduleRepo:  s,
	}
}

// === Client Actions ===

func (uc *BookingInteractor) ListAvailableSlots(procedureID, date string) ([]domain.TimeSlot, error) {
	return uc.ScheduleRepo.GetAvailableSlots(procedureID, date)
}

func (uc *BookingInteractor) CreateBooking(clientID, specialistID, procedureID string, start time.Time) (*domain.Booking, error) {
	proc, err := uc.ProcedureRepo.GetByID(procedureID)
	if err != nil {
		return nil, errors.New("procedure not found")
	}
	end := start.Add(time.Duration(proc.DurationMinutes) * time.Minute)

	isFree, err := uc.BookingRepo.IsSlotAvailable(specialistID, start.Format(time.RFC3339), end.Format(time.RFC3339))
	if err != nil || !isFree {
		return nil, errors.New("time slot unavailable")
	}

	booking := &domain.Booking{
		ClientID:     clientID,
		SpecialistID: specialistID,
		ProcedureID:  procedureID,
		StartTime:    start,
		EndTime:      end,
		Status:       "booked",
		CreatedAt:    time.Now(),
	}
	return uc.BookingRepo.Create(booking)
}

func (uc *BookingInteractor) CancelBooking(id string) error {
	return uc.BookingRepo.Cancel(id)
}

func (uc *BookingInteractor) GetClientBookings(clientID string) ([]*domain.Booking, error) {
	return uc.BookingRepo.GetByClient(clientID)
}

// === Admin Actions ===

func (uc *BookingInteractor) CreateProcedure(name string, duration int32) (*domain.Procedure, error) {
	if name == "" || duration <= 0 {
		return nil, errors.New("invalid procedure input")
	}
	p := &domain.Procedure{
		Name:            name,
		DurationMinutes: duration,
		CreatedAt:       time.Now(),
	}
	return uc.ProcedureRepo.Create(p)
}

func (uc *BookingInteractor) AssignProcedure(specialistID, procedureID string) error {
	return uc.ScheduleRepo.AssignProcedure(specialistID, procedureID)
}

func (uc *BookingInteractor) CreateScheduleTemplate(specialistID string, weekday int, start, end string, breakMins int32) error {
	template := &domain.ScheduleTemplate{
		SpecialistID: specialistID,
		Weekday:      weekday,
		StartTime:    start,
		EndTime:      end,
		BreakMinutes: breakMins,
	}
	return uc.ScheduleRepo.CreateTemplate(template)
}

func (uc *BookingInteractor) OverrideDaySchedule(specialistID, date, start, end string, cancelled bool) error {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return fmt.Errorf("invalid date format")
	}
	override := &domain.DailySchedule{
		SpecialistID: specialistID,
		Date:         t,
		StartTime:    start,
		EndTime:      end,
		Override:     true,
		Cancelled:    cancelled,
	}
	return uc.ScheduleRepo.OverrideDailySchedule(override)
}
