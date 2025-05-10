package usecase

import (
	"errors"

	"github.com/gemdivk/LUMERA-SPA/user-service/internal/domain"
	"github.com/gemdivk/LUMERA-SPA/user-service/internal/domain/repository"
	"github.com/gemdivk/LUMERA-SPA/user-service/internal/infrastructure/auth"
	"golang.org/x/crypto/bcrypt"
)

type UserInteractor struct {
	Repo repository.UserRepo
}

func NewUserInteractor(repo repository.UserRepo) *UserInteractor {
	return &UserInteractor{Repo: repo}
}

func (u *UserInteractor) Register(name, email, password string) (*domain.User, string, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 12)

	user, err := u.Repo.Create(name, email, string(hash))
	if err != nil {
		return nil, "", err
	}
	_ = u.Repo.AssignRole(user.ID, "client") // 2 = client

	count, _ := u.Repo.CountUsers()
	if count == 1 {
		_ = u.Repo.AssignRole(user.ID, "admin") // 1 = admin
	}

	roles, _ := u.Repo.GetRoles(user.ID)
	token, err := auth.GenerateToken(user.ID, user.Email, roles)
	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}

func (u *UserInteractor) Login(email, password string) (string, error) {
	user, err := u.Repo.GetByEmail(email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	roles, _ := u.Repo.GetRoles(user.ID)
	return auth.GenerateToken(user.ID, user.Email, roles)
}

func (u *UserInteractor) GetProfile(userID string) (*domain.User, error) {
	return u.Repo.GetByID(userID)
}

func (u *UserInteractor) GetProfileFromEmail(email string) (*domain.User, error) {
	return u.Repo.GetByEmail(email)
}

func (u *UserInteractor) UpdateProfile(id, name, password string) (*domain.User, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	return u.Repo.UpdateNameAndPassword(id, name, string(hash))
}

func (u *UserInteractor) AssignRole(userID string, roleName string) error {
	return u.Repo.AssignRole(userID, roleName)
}

func (u *UserInteractor) ListRoles(userID string) ([]string, error) {
	return u.Repo.GetRoles(userID)
}

func (u *UserInteractor) GetAllUsers() ([]*domain.User, error) {
	return u.Repo.GetAll()
}

func (u *UserInteractor) SearchUsers(query string) ([]*domain.User, error) {
	return u.Repo.Search(query)
}

func (u *UserInteractor) DeleteUser(userID string) error {
	return u.Repo.Delete(userID)
}

func (u *UserInteractor) RemoveRole(userID string, roleName string) error {
	return u.Repo.RemoveRole(userID, roleName)
}
