package model

type UserAnswer struct {
	Model
	Content    string `json:"content" db:"content"`
	QuestionID int    `json:"question_id" db:"question_id"`
	IsTrue     bool   `json:"is_true" db:"is_true"`
	IsChosen   bool   `json:"is_chosen" db:"is_chosen"`
	Sort       int    `json:"sort" db:"sort"`
}
