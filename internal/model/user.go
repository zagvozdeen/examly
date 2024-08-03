package model

import (
	"github.com/guregu/null/v5"
	"time"
)

type User struct {
	ID        int         `json:"id" db:"id"`
	Email     null.String `json:"email" db:"email"`
	FirstName null.String `json:"first_name" db:"first_name"`
	LastName  null.String `json:"last_name" db:"last_name"`
	Role      UserRole    `json:"role" db:"role"`
	Password  null.String `json:"password" db:"password"`
	AvatarID  null.Int    `json:"avatar_id" db:"avatar_id"`
	DeletedAt null.Time   `json:"deleted_at" db:"deleted_at"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" db:"updated_at"`
}
