package store

import (
	"context"
	"errors"
	"github.com/guregu/null/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zagvozdeen/examly/internal/enum"
	"time"
)

type User struct {
	ID        int           `json:"id"`
	UUID      string        `json:"uuid"`
	Email     null.String   `json:"email"`
	FirstName null.String   `json:"first_name"`
	LastName  null.String   `json:"last_name"`
	FullName  null.String   `json:"full_name"`
	Role      enum.UserRole `json:"role"`
	Password  null.String   `json:"-"`
	AvatarID  null.Int      `json:"avatar_id"`
	DeletedAt null.Time     `json:"deleted_at"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type UsersStore interface {
	Get(ctx context.Context) ([]User, error)
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByID(ctx context.Context, id int) (User, error)
	Update(ctx context.Context, user *User) error
}

type UserStore struct {
	conn *pgxpool.Pool
}

func (s *UserStore) Get(ctx context.Context) (users []User, err error) {
	sql := "SELECT id, uuid, email, first_name, last_name, role, avatar_id, deleted_at, created_at, updated_at FROM users WHERE deleted_at IS NULL ORDER BY created_at DESC"
	rows, err := s.conn.Query(ctx, sql)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return users, ErrNotFound
		}
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(
			&user.ID,
			&user.UUID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Role,
			&user.AvatarID,
			&user.DeletedAt,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return
		}
		users = append(users, user)
	}

	return
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	return s.conn.QueryRow(
		ctx,
		`INSERT INTO users
				(uuid, email, first_name, last_name, role, password, avatar_id, created_at, updated_at)
			VALUES 
			    ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			RETURNING id`,
		user.UUID,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Role,
		user.Password,
		user.AvatarID,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID)
}

func (s *UserStore) GetByEmail(ctx context.Context, email string) (user User, err error) {
	err = s.conn.QueryRow(
		ctx,
		"SELECT id, password FROM users WHERE email = $1 AND deleted_at IS NULL",
		email,
	).Scan(
		&user.ID,
		&user.Password,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return user, ErrNotFound
	}

	return
}

func (s *UserStore) GetByID(ctx context.Context, id int) (user User, err error) {
	err = s.conn.QueryRow(
		ctx,
		"SELECT id, uuid, email, first_name, last_name, role, password, avatar_id, deleted_at, created_at, updated_at FROM users WHERE id = $1 AND deleted_at IS NULL",
		id,
	).Scan(
		&user.ID,
		&user.UUID,
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

	if errors.Is(err, pgx.ErrNoRows) {
		return user, ErrNotFound
	}

	return
}

func (s *UserStore) Update(ctx context.Context, user *User) error {
	_, err := s.conn.Exec(
		ctx,
		`UPDATE users
			SET email = $1, first_name = $2, last_name = $3, updated_at = $4
			WHERE id = $5 AND deleted_at IS NULL`,
		user.Email,
		user.FirstName,
		user.LastName,
		user.UpdatedAt,
		user.ID,
	)
	return err
}
