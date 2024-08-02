package repository

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/jmoiron/sqlx"
)

type CourseRepository struct {
	db *sqlx.DB
}

func NewCourseRepository(db *sqlx.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (r *CourseRepository) GetCourses() ([]model.Course, error) {
	var courses []model.Course

	err := r.db.Select(
		&courses,
		"SELECT * FROM courses WHERE deleted_at IS NULL",
	)

	return courses, err
}
