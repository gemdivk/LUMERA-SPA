package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
	"testing"
	"time"

	"github.com/gemdivk/LUMERA-SPA/review-service/internal/domain"
	"github.com/gemdivk/LUMERA-SPA/review-service/internal/infrastructure/postgres"
	"github.com/stretchr/testify/assert"
)

func TestCreateReview(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewReviewRepo(db)
	review := &domain.Review{
		SalonID: "salon-123",
		UserID:  "user-123",
		Content: "Great service!",
		Rating:  5,
	}

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO reviews (id, salon_id, user_id, content, rating, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)`)).
		WithArgs(sqlmock.AnyArg(), review.SalonID, review.UserID, review.Content, review.Rating, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := repo.Create(review)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.ID)
	assert.Equal(t, review.Content, res.Content)
}

func TestGetReview(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewReviewRepo(db)
	reviewID := "test-id"
	expected := &domain.Review{
		ID:        reviewID,
		SalonID:   "salon-123",
		UserID:    "user-123",
		Content:   "Amazing",
		Rating:    4,
		CreatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "salon_id", "user_id", "content", "rating", "created_at"}).
		AddRow(expected.ID, expected.SalonID, expected.UserID, expected.Content, expected.Rating, expected.CreatedAt)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, salon_id, user_id, content, rating, created_at FROM reviews WHERE id=$1`)).
		WithArgs(reviewID).
		WillReturnRows(rows)

	got, err := repo.Get(reviewID)
	assert.NoError(t, err)
	assert.Equal(t, expected.Content, got.Content)
}

func TestUpdateReview(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewReviewRepo(db)
	review := &domain.Review{
		ID:      "test-id",
		Content: "Updated review",
		Rating:  3,
	}

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE reviews SET content=$1, rating=$2 WHERE id=$3`)).
		WithArgs(review.Content, review.Rating, review.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := repo.Update(review)
	assert.NoError(t, err)
	assert.Equal(t, "Updated review", res.Content)
}

func TestDeleteReview(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewReviewRepo(db)
	reviewID := "test-id"

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM reviews WHERE id=$1`)).
		WithArgs(reviewID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Delete(reviewID)
	assert.NoError(t, err)
}

func TestListBySalon(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewReviewRepo(db)
	salonID := "salon-123"
	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "salon_id", "user_id", "content", "rating", "created_at"}).
		AddRow("1", salonID, "user1", "Nice!", 5, now).
		AddRow("2", salonID, "user2", "Average", 3, now)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, salon_id, user_id, content, rating, created_at FROM reviews WHERE salon_id=$1`)).
		WithArgs(salonID).
		WillReturnRows(rows)

	reviews, err := repo.ListBySalon(salonID)
	assert.NoError(t, err)
	assert.Len(t, reviews, 2)
	assert.Equal(t, "Nice!", reviews[0].Content)
}
