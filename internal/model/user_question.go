package model

import "github.com/guregu/null/v5"

type UserQuestion struct {
	Model
	UUID        string       `json:"uuid" db:"uuid"`
	Content     string       `json:"content" db:"content"`
	Explanation null.String  `json:"explanation" db:"explanation"`
	Type        string       `json:"type" db:"type"`
	IsTrue      null.Bool    `json:"is_true" db:"is_true"`
	Sort        int          `json:"sort" db:"sort"`
	CourseID    int          `json:"course_id" db:"course_id"`
	QuestionID  int          `json:"question_id" db:"question_id"`
	ModuleID    null.Int     `json:"module_id" db:"module_id"`
	FileID      null.Int     `json:"file_id" db:"file_id"`
	Answers     []UserAnswer `json:"answers"`
}
