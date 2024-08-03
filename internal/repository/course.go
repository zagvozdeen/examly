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

func (r *CourseRepository) CreateCourse(course *model.Course) (id int, err error) {
	err = r.db.QueryRow(
		"INSERT INTO courses (uuid, name, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		course.UUID,
		course.Name,
		course.UserID,
		course.CreatedAt,
		course.UpdatedAt,
	).Scan(&id)

	return id, err
}
