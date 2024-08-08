package model

import "github.com/guregu/null/v5"

type UserCourse struct {
	Model
	UUID           string         `json:"uuid" db:"uuid"`
	Name           string         `json:"name" db:"name"`
	Type           string         `json:"type" db:"type"`
	UserID         int            `json:"user_id" db:"user_id"`
	CourseID       int            `json:"course_id" db:"course_id"`
	LastQuestionID null.Int       `json:"last_question_id" db:"last_question_id"`
	Questions      []UserQuestion `json:"questions"`
	Modules        []UserModule   `json:"modules"`
}
