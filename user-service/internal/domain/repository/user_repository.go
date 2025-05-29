package repository

import "github.com/gemdivk/LUMERA-SPA/user-service/internal/domain"

type UserRepo interface {
	Create(name, email, password string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	GetByID(id string) (*domain.User, error)
	UpdateNameAndPassword(id, name, password string) (*domain.User, error)
	AssignRole(userID string, roleName string) error
	RemoveRole(userID string, roleName string) error
	GetRoles(userID string) ([]string, error)
	GetAll() ([]*domain.User, error)
	Search(query string) ([]*domain.User, error)
	Delete(userID string) error
	CountUsers() (int, error)
	MarkEmailVerified(userID string) error
}
