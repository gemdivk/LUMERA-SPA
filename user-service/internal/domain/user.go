package domain

import "time"

type User struct {
	ID         string
	Name       string
	Email      string
	Password   string
	CreatedAt  time.Time
	IsVerified bool
}

type Role struct {
	ID   int
	Name string
}
