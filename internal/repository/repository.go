package repository

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/jmoiron/sqlx"
)

type Courses interface {
	GetCourses() ([]model.Course, error)
}

type Repository struct {
	Courses
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Courses: NewCourseRepository(db),
	}
}
