package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int            `db:"id"`
	Email     sql.NullString `db:"email"`
	FirstName sql.NullString `db:"first_name"`
	LastName  sql.NullString `db:"last_name"`
	Role      UserRole       `db:"role"`
	Password  sql.NullString `db:"password"`
	AvatarID  sql.NullInt64  `db:"avatar_id"`
	DeletedAt sql.NullTime   `db:"deleted_at"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}
