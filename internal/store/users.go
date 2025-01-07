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
	ID                int           `json:"id"`
	UUID              string        `json:"uuid"`
	Email             null.String   `json:"email"`
	FirstName         null.String   `json:"first_name"`
	LastName          null.String   `json:"last_name"`
	FullName          null.String   `json:"full_name"`
	Role              enum.UserRole `json:"role"`
	Password          null.String   `json:"-"`
	AvatarID          null.Int      `json:"avatar_id"`
	Description       null.String   `json:"description"`
	CompanyName       null.String   `json:"company_name"`
	Contact           null.String   `json:"contact"`
	Account           int           `json:"account"`
	CanViewReferrals  bool          `json:"can_view_referrals"`
	DeletedAt         null.Time     `json:"deleted_at"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
	HasUserExperience bool          `json:"has_user_experience"`
}

type UserExperience struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	One       int       `json:"one"`
	Two       int       `json:"two"`
	Three     int       `json:"three"`
	Four      string    `json:"four"`
	Five      int       `json:"five"`
	Six       int       `json:"six"`
	Seven     string    `json:"seven"`
	Eight     string    `json:"eight"`
	Nine      int       `json:"nine"`
	Ten       string    `json:"ten"`
	Eleven    int       `json:"eleven"`
	Twelve    string    `json:"twelve"`
	Thirteen  string    `json:"thirteen"`
	DeletedAt null.Time `json:"deleted_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UsersStore interface {
	Get(ctx context.Context) ([]User, error)
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByID(ctx context.Context, id int) (User, error)
	Update(ctx context.Context, user *User) error
	UpdateAccount(ctx context.Context, user *User) error
	UpdateCanViewReferrals(ctx context.Context, user *User) error
	GetUserExperience(ctx context.Context, id int) (UserExperience, error)
	CreateUserExperience(ctx context.Context, ue *UserExperience) error
	GetReferrals(ctx context.Context, id int) (users []User, err error)
}

type UserStore struct {
	conn *pgxpool.Pool
}

func (s *UserStore) Get(ctx context.Context) (users []User, err error) {
	sql := "SELECT id, uuid, email, first_name, last_name, role, avatar_id, deleted_at, created_at, updated_at, contact FROM users WHERE deleted_at IS NULL ORDER BY created_at DESC"
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
			&user.Contact,
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
	).Scan(&user.ID, &user.Password)
	if errors.Is(err, pgx.ErrNoRows) {
		err = ErrNotFound
	}
	return
}

func (s *UserStore) GetByID(ctx context.Context, id int) (user User, err error) {
	err = s.conn.QueryRow(
		ctx,
		"SELECT id, uuid, email, first_name, last_name, role, password, avatar_id, deleted_at, created_at, updated_at, description, company_name, contact, account, can_view_referrals FROM users WHERE id = $1 AND deleted_at IS NULL",
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
		&user.Description,
		&user.CompanyName,
		&user.Contact,
		&user.Account,
		&user.CanViewReferrals,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		err = ErrNotFound
	}
	return
}

func (s *UserStore) Update(ctx context.Context, user *User) error {
	_, err := s.conn.Exec(
		ctx,
		"UPDATE users SET email = $1, first_name = $2, last_name = $3, description = $4, company_name = $5, contact = $6, avatar_id = $7, updated_at = $8, role = $9 WHERE id = $10",
		user.Email,
		user.FirstName,
		user.LastName,
		user.Description,
		user.CompanyName,
		user.Contact,
		user.AvatarID,
		user.UpdatedAt,
		user.Role,
		user.ID,
	)
	return err
}

func (s *UserStore) UpdateAccount(ctx context.Context, user *User) error {
	_, err := s.conn.Exec(
		ctx,
		"UPDATE users SET account = $1, updated_at = $2 WHERE id = $3",
		user.Account,
		user.UpdatedAt,
		user.ID,
	)
	return err
}

func (s *UserStore) UpdateCanViewReferrals(ctx context.Context, user *User) error {
	_, err := s.conn.Exec(
		ctx,
		"UPDATE users SET can_view_referrals = $1, account = $2, updated_at = $3 WHERE id = $4",
		user.CanViewReferrals,
		user.Account,
		user.UpdatedAt,
		user.ID,
	)
	return err
}

func (s *UserStore) GetUserExperience(ctx context.Context, id int) (ue UserExperience, err error) {
	err = s.conn.QueryRow(
		ctx,
		"SELECT id, user_id, one, two, three, four, five, six, seven, eight, nine, ten, eleven, twelve, thirteen, deleted_at, created_at, updated_at FROM user_experience WHERE user_id = $1 AND deleted_at IS NULL",
		id,
	).Scan(
		&ue.ID,
		&ue.UserID,
		&ue.One,
		&ue.Two,
		&ue.Three,
		&ue.Four,
		&ue.Five,
		&ue.Six,
		&ue.Seven,
		&ue.Eight,
		&ue.Nine,
		&ue.Ten,
		&ue.Eleven,
		&ue.Twelve,
		&ue.Thirteen,
		&ue.DeletedAt,
		&ue.CreatedAt,
		&ue.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		err = ErrNotFound
	}
	return
}

func (s *UserStore) CreateUserExperience(ctx context.Context, ue *UserExperience) error {
	return s.conn.QueryRow(
		ctx,
		"INSERT INTO user_experience (user_id, one, two, three, four, five, six, seven, eight, nine, ten, eleven, twelve, thirteen, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING id",
		ue.UserID,
		ue.One,
		ue.Two,
		ue.Three,
		ue.Four,
		ue.Five,
		ue.Six,
		ue.Seven,
		ue.Eight,
		ue.Nine,
		ue.Ten,
		ue.Eleven,
		ue.Twelve,
		ue.Thirteen,
		ue.CreatedAt,
		ue.UpdatedAt,
	).Scan(&ue.ID)
}

func (s *UserStore) GetReferrals(ctx context.Context, id int) (users []User, err error) {
	rows, err := s.conn.Query(
		ctx,
		"SELECT id, uuid, email, first_name, last_name, role, avatar_id, deleted_at, created_at, updated_at, contact FROM users WHERE id != $1 AND role IN ($2, $3) AND deleted_at IS NULL ORDER BY created_at DESC",
		id,
		enum.CompanyRole.String(),
		enum.ReferralRole.String(),
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = ErrNotFound
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
			&user.Contact,
		)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	return
}
