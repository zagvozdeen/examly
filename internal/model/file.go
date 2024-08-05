package model

import (
	"github.com/guregu/null/v5"
	"time"
)

type File struct {
	ID         int       `json:"id" db:"id"`
	UUID       string    `json:"uuid" db:"uuid"`
	Content    string    `json:"content" db:"content"`
	Size       int       `json:"size" db:"size"`
	MimeType   string    `json:"mime_type" db:"mime_type"`
	OriginName string    `json:"origin_name" db:"origin_name"`
	UserID     int       `json:"user_id" db:"user_id"`
	DeletedAt  null.Time `json:"deleted_at" db:"deleted_at"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
