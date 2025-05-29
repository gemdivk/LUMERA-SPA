package usecase

import (
	"errors"
	"testing"

	"github.com/gemdivk/LUMERA-SPA/user-service/internal/domain"
	"github.com/gemdivk/LUMERA-SPA/user-service/internal/infrastructure/cache"
	"github.com/stretchr/testify/assert"
)

type MockRepo struct {
	users   map[string]*domain.User
	roles   map[string][]string
	created bool
}

func (m *MockRepo) Create(name, email, password string) (*domain.User, error) {
	user := &domain.User{ID: "u3", Name: name, Email: email, Password: password, IsVerified: false}
	m.users[user.ID] = user
	m.created = true
	return user, nil
}
func (m *MockRepo) GetByEmail(email string) (*domain.User, error) {
	for _, u := range m.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, errors.New("not found")
}
func (m *MockRepo) GetByID(id string) (*domain.User, error) {
	u, ok := m.users[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return u, nil
}
func (m *MockRepo) UpdateNameAndPassword(id, name, password string) (*domain.User, error) {
	u, ok := m.users[id]
	if !ok {
		return nil, errors.New("not found")
	}
	u.Name = name
	u.Password = password
	return u, nil
}
func (m *MockRepo) AssignRole(userID string, roleName string) error {
	m.roles[userID] = append(m.roles[userID], roleName)
	return nil
}
func (m *MockRepo) RemoveRole(userID string, roleName string) error {
	roles := m.roles[userID]
	var updated []string
	for _, r := range roles {
		if r != roleName {
			updated = append(updated, r)
		}
	}
	m.roles[userID] = updated
	return nil
}
func (m *MockRepo) GetRoles(userID string) ([]string, error) {
	return m.roles[userID], nil
}
func (m *MockRepo) GetAll() ([]*domain.User, error) {
	var list []*domain.User
	for _, u := range m.users {
		list = append(list, u)
	}
	return list, nil
}
func (m *MockRepo) Search(query string) ([]*domain.User, error) {
	var result []*domain.User
	for _, u := range m.users {
		if u.Name == query || u.Email == query {
			result = append(result, u)
		}
	}
	return result, nil
}
func (m *MockRepo) Delete(userID string) error {
	delete(m.users, userID)
	return nil
}
func (m *MockRepo) CountUsers() (int, error) {
	return len(m.users), nil
}
func (m *MockRepo) MarkEmailVerified(userID string) error {
	if u, ok := m.users[userID]; ok {
		u.IsVerified = true
		return nil
	}
	return errors.New("not found")
}

func setup() (*UserInteractor, *MockRepo) {
	repo := &MockRepo{
		users: map[string]*domain.User{
			"u1": {ID: "u1", Name: "John", Email: "john@example.com", Password: "$2a$12$LkzNkwJrUeFbFa9Tx8u69.T9q8oz1ceobq54FrIXChMiR5PUKxG/S", IsVerified: true},
			"u2": {ID: "u2", Name: "Unverified", Email: "unverified@example.com", Password: "$2a$12$LkzNkwJrUeFbFa9Tx8u69.T9q8oz1ceobq54FrIXChMiR5PUKxG/S", IsVerified: false},
		},
		roles: map[string][]string{
			"u1": {"client"},
		},
	}
	cache := cache.NewUserCache()
	cache.LoadInitial([]*domain.User{repo.users["u1"], repo.users["u2"]})
	u := NewUserInteractorWithCache(repo, nil, cache)
	return u, repo
}

func TestUpdateProfile(t *testing.T) {
	u, _ := setup()
	user, err := u.UpdateProfile("u1", "Updated", "newpass")
	assert.NoError(t, err)
	assert.Equal(t, "Updated", user.Name)
}

func TestAssignAndRemoveRole(t *testing.T) {
	u, repo := setup()
	err := u.AssignRole("u1", "admin")
	assert.NoError(t, err)
	assert.Contains(t, repo.roles["u1"], "admin")

	err = u.RemoveRole("u1", "admin")
	assert.NoError(t, err)
	assert.NotContains(t, repo.roles["u1"], "admin")
}

func TestListRoles(t *testing.T) {
	u, _ := setup()
	roles, err := u.ListRoles("u1")
	assert.NoError(t, err)
	assert.Equal(t, []string{"client"}, roles)
}

func TestGetAllUsers(t *testing.T) {
	u, _ := setup()
	users, err := u.GetAllUsers()
	assert.NoError(t, err)
	assert.Len(t, users, 2)
}

func TestSearchUsers(t *testing.T) {
	u, _ := setup()
	found, err := u.SearchUsers("john@example.com")
	assert.NoError(t, err)
	assert.Len(t, found, 1)
	assert.Equal(t, "u1", found[0].ID)
}

func TestDeleteUser(t *testing.T) {
	u, repo := setup()
	err := u.DeleteUser("u2")
	assert.NoError(t, err)
	_, ok := repo.users["u2"]
	assert.False(t, ok)
}

func TestGetProfileFromEmail(t *testing.T) {
	u, _ := setup()
	user, err := u.GetProfileFromEmail("john@example.com")
	assert.NoError(t, err)
	assert.Equal(t, "u1", user.ID)
}
