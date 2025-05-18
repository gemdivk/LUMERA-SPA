package postgres

import (
	"database/sql"
	"time"

	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain"
	"github.com/google/uuid"
)

type BookingRepo struct {
	DB *sql.DB
}

func NewBookingRepo(db *sql.DB) *BookingRepo {
	return &BookingRepo{DB: db}
}

func (r *BookingRepo) Create(b *domain.Booking) (*domain.Booking, error) {
	b.ID = uuid.New().String()
	b.CreatedAt = time.Now()
	_, err := r.DB.Exec(`
		INSERT INTO bookings (id, client_id, specialist_id, procedure_id, start_time, end_time, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`, b.ID, b.ClientID, b.SpecialistID, b.ProcedureID, b.StartTime, b.EndTime, b.Status, b.CreatedAt)
	return b, err
}

func (r *BookingRepo) Cancel(id string) error {
	_, err := r.DB.Exec(`UPDATE bookings SET status='cancelled' WHERE id=$1`, id)
	return err
}

func (r *BookingRepo) GetByClient(clientID string) ([]*domain.Booking, error) {
	rows, err := r.DB.Query(`SELECT id, specialist_id, procedure_id, start_time, end_time, status, created_at FROM bookings WHERE client_id=$1`, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*domain.Booking
	for rows.Next() {
		b := &domain.Booking{ClientID: clientID}
		err := rows.Scan(&b.ID, &b.SpecialistID, &b.ProcedureID, &b.StartTime, &b.EndTime, &b.Status, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, b)
	}
	return result, nil
}

func (r *BookingRepo) IsSlotAvailable(specialistID, start, end string) (bool, error) {
	var count int
	err := r.DB.QueryRow(`
		SELECT COUNT(*) FROM bookings 
		WHERE specialist_id = $1 AND status = 'booked'
		AND start_time < $3 AND end_time > $2
	`, specialistID, start, end).Scan(&count)
	return count == 0, err
}
