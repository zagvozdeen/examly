package model

import (
	"github.com/guregu/null/v5"
	"time"
)

type Answer struct {
	ID         int       `json:"id" db:"id"`
	Content    string    `json:"content" db:"content"`
	QuestionID int       `json:"question_id" db:"question_id"`
	IsTrue     bool      `json:"is_true" db:"is_true"`
	DeletedAt  null.Time `json:"deleted_at" db:"deleted_at"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
