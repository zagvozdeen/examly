package model

import (
	"github.com/guregu/null/v5"
	"time"
)

type Question struct {
	ID        int       `json:"id" db:"id"`
	Content   string    `json:"content" db:"content"`
	Type      string    `json:"type" db:"type"`
	CourseID  int       `json:"course_id" db:"course_id"`
	ModuleID  null.Int  `json:"module_id" db:"module_id"`
	DeletedAt null.Time `json:"deleted_at" db:"deleted_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (q *Question) GetType() QuestionType {
	return QuestionType(q.Type)
}
