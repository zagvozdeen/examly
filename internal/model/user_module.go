package model

type UserModule struct {
	Model
	Name     string `json:"name" db:"name"`
	CourseID int    `json:"course_id" db:"course_id"`
}
