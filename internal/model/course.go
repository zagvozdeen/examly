package model

import (
	"database/sql"
	"time"
)

type Course struct {
	ID        int          `db:"id"`
	UUID      string       `db:"uuid"`
	Name      string       `db:"name"`
	UserID    int          `db:"user_id"`
	DeletedAt sql.NullTime `db:"deleted_at"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
}
