package domain

import "time"

type EmailLog struct {
	ID      int
	Email   string
	Subject string
	SentAt  time.Time
}
