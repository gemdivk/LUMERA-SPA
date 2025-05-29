package postgres

import (
	"database/sql"
	"github.com/gemdivk/LUMERA-SPA/notification-service/internal/domain"
	"time"
)

type EmailLogRepo struct {
	DB *sql.DB
}

func NewEmailLogRepo(db *sql.DB) *EmailLogRepo {
	return &EmailLogRepo{DB: db}
}

func (r *EmailLogRepo) LogEmail(email, subject string) error {
	_, err := r.DB.Exec(`INSERT INTO email_logs (email, subject, sent_at) VALUES ($1, $2, $3)`,
		email, subject, time.Now())
	return err
}

func (r *EmailLogRepo) FetchAll() ([]domain.EmailLog, error) {
	rows, err := r.DB.Query(`SELECT id, email, subject, sent_at FROM email_logs ORDER BY sent_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []domain.EmailLog
	for rows.Next() {
		var log domain.EmailLog
		if err := rows.Scan(&log.ID, &log.Email, &log.Subject, &log.SentAt); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}
