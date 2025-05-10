package application

import "github.com/gemdivk/LUMERA-SPA/user-service/internal/domain"

type UserUsecase interface {
	Register(name, email, password string) (*domain.User, string, error)
	Login(email, password string) (string, error)
	GetProfile(userID string) (*domain.User, error)
	GetProfileFromEmail(email string) (*domain.User, error)
	UpdateProfile(userID, name, password string) (*domain.User, error)
	AssignRole(userID string, roleName string) error
	RemoveRole(userID string, roleName string) error
	ListRoles(userID string) ([]string, error)
	GetAllUsers() ([]*domain.User, error)
	SearchUsers(query string) ([]*domain.User, error)
	DeleteUser(userID string) error
}
