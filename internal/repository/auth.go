package repository

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user model.User) (int, error) {
	var id int

	err := r.db.QueryRow(
		"INSERT INTO users (email, first_name, last_name, role, password, avatar_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		user.Email,
		user.FirstName,
		user.LastName,
		user.Role,
		user.Password,
		user.AvatarID,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&id)

	return id, err
}

func (r *AuthRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User

	row := r.db.QueryRow(
		"SELECT id, email, first_name, last_name, role, password, avatar_id, deleted_at, created_at, updated_at FROM users WHERE email = $1 AND deleted_at IS NULL LIMIT 1",
		email,
	)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Role,
		&user.Password,
		&user.AvatarID,
		&user.DeletedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, err
}

func (r *AuthRepository) IsExistsUserByEmail(email string) (bool, error) {
	var exists bool

	err := r.db.Get(
		&exists,
		"SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 AND deleted_at IS NULL)",
		email,
	)

	return exists, err
}

func (r *AuthRepository) GetUserByID(id int) (user model.User, err error) {
	row := r.db.QueryRow(
		"SELECT id, email, first_name, last_name, nullif(concat_ws(' ', first_name, last_name), '') as full_name, role, password, avatar_id, deleted_at, created_at, updated_at FROM users WHERE id = $1 AND deleted_at IS NULL LIMIT 1",
		id,
	)

	err = row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.FullName,
		&user.Role,
		&user.Password,
		&user.AvatarID,
		&user.DeletedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, err
}

func (r *AuthRepository) UpdateUser(user *model.User) error {
	_, err := r.db.Exec(
		"UPDATE users SET first_name = $1, last_name = $2, email = $3, updated_at = $4 WHERE id = $5 AND deleted_at IS NULL",
		user.FirstName,
		user.LastName,
		user.Email,
		user.UpdatedAt,
		user.ID,
	)

	return err
}
