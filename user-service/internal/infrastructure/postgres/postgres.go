package postgres

import (
	"database/sql"
	"time"

	"github.com/gemdivk/LUMERA-SPA/user-service/internal/domain"
	"github.com/google/uuid"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Create(name, email, password string) (*domain.User, error) {
	id := uuid.New().String()
	created := time.Now()
	_, err := r.DB.Exec(`INSERT INTO auth_users (id, name, email, password, created_at, is_verified) VALUES ($1, $2, $3, $4, $5, $6)`,
		id, name, email, password, created, false)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:         id,
		Name:       name,
		Email:      email,
		Password:   password,
		CreatedAt:  created,
		IsVerified: false,
	}, nil
}

func (r *UserRepo) GetByEmail(email string) (*domain.User, error) {
	row := r.DB.QueryRow(`SELECT id, name, email, password, created_at, is_verified FROM auth_users WHERE email=$1`, email)
	u := &domain.User{}
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt, &u.IsVerified)
	return u, err
}

func (r *UserRepo) GetByID(id string) (*domain.User, error) {
	row := r.DB.QueryRow(`SELECT id, name, email, password, created_at, is_verified FROM auth_users WHERE id=$1`, id)
	u := &domain.User{}
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt, &u.IsVerified)
	return u, err
}

func (r *UserRepo) AssignRole(userID string, roleName string) error {
	var roleID int
	err := r.DB.QueryRow(`SELECT id FROM auth_roles WHERE name = $1`, roleName).Scan(&roleID)
	if err != nil {
		return err
	}
	_, err = r.DB.Exec(`INSERT INTO auth_user_roles (user_id, role_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`, userID, roleID)
	return err
}

func (r *UserRepo) GetRoles(userID string) ([]string, error) {
	rows, err := r.DB.Query(`
		SELECT name FROM auth_roles
		JOIN auth_user_roles ON auth_roles.id = auth_user_roles.role_id
		WHERE auth_user_roles.user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		roles = append(roles, name)
	}
	return roles, nil
}

func (r *UserRepo) UpdateNameAndPassword(id, name, password string) (*domain.User, error) {
	_, err := r.DB.Exec(`UPDATE auth_users SET name=$1, password=$2 WHERE id=$3`, name, password, id)
	if err != nil {
		return nil, err
	}
	return r.GetByID(id)
}
func (r *UserRepo) GetAll() ([]*domain.User, error) {
	rows, err := r.DB.Query(`SELECT id, name, email, password, created_at FROM auth_users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		u := &domain.User{}
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepo) Search(query string) ([]*domain.User, error) {
	rows, err := r.DB.Query(`SELECT id, name, email, password, created_at FROM auth_users WHERE name ILIKE '%' || $1 || '%' OR email ILIKE '%' || $1 || '%'`, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		u := &domain.User{}
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepo) Delete(userID string) error {
	_, err := r.DB.Exec(`DELETE FROM auth_users WHERE id = $1`, userID)
	return err
}
func (r *UserRepo) CountUsers() (int, error) {
	row := r.DB.QueryRow(`SELECT COUNT(*) FROM auth_users`)
	var count int
	err := row.Scan(&count)
	return count, err
}

func (r *UserRepo) RemoveRole(userID string, roleName string) error {
	var roleID int
	err := r.DB.QueryRow(`SELECT id FROM auth_roles WHERE name = $1`, roleName).Scan(&roleID)
	if err != nil {
		return err
	}
	_, err = r.DB.Exec(`DELETE FROM auth_user_roles WHERE user_id = $1 AND role_id = $2`, userID, roleID)
	return err
}
func (r *UserRepo) MarkEmailVerified(userID string) error {
	_, err := r.DB.Exec(`UPDATE auth_users SET is_verified = true WHERE id = $1`, userID)
	return err
}
