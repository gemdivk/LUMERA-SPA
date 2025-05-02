package postgres

import (
	"database/sql"
	"github.com/gemdivk/LUMERA-SPA/review-service/internal/domain"
	"time"

	"github.com/google/uuid"
)

type ReviewRepo struct {
	DB *sql.DB
}

func NewReviewRepo(db *sql.DB) *ReviewRepo {
	return &ReviewRepo{DB: db}
}
func (r *ReviewRepo) Create(review *domain.Review) (*domain.Review, error) {
	review.ID = uuid.New().String()
	review.CreatedAt = time.Now()
	_, err := r.DB.Exec(`INSERT INTO reviews (id, salon_id, user_id, content, rating, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		review.ID, review.SalonID, review.UserID, review.Content, review.Rating, review.CreatedAt)
	return review, err
}

func (r *ReviewRepo) Get(id string) (*domain.Review, error) {
	row := r.DB.QueryRow(`SELECT id, salon_id, user_id, content, rating, created_at FROM reviews WHERE id=$1`, id)
	review := &domain.Review{}
	err := row.Scan(&review.ID, &review.SalonID, &review.UserID, &review.Content, &review.Rating, &review.CreatedAt)
	return review, err
}

func (r *ReviewRepo) Update(review *domain.Review) (*domain.Review, error) {
	_, err := r.DB.Exec(`UPDATE reviews SET content=$1, rating=$2 WHERE id=$3`,
		review.Content, review.Rating, review.ID)
	return review, err
}

func (r *ReviewRepo) Delete(id string) error {
	_, err := r.DB.Exec(`DELETE FROM reviews WHERE id=$1`, id)
	return err
}

func (r *ReviewRepo) ListBySalon(salonID string) ([]*domain.Review, error) {
	rows, err := r.DB.Query(`SELECT id, salon_id, user_id, content, rating, created_at FROM reviews WHERE salon_id=$1`, salonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []*domain.Review
	for rows.Next() {
		r := &domain.Review{}
		err := rows.Scan(&r.ID, &r.SalonID, &r.UserID, &r.Content, &r.Rating, &r.CreatedAt)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, r)
	}
	return reviews, nil
}
