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
	courses := make([]model.Course, 0)

	err := r.db.Select(
		&courses,
		"SELECT * FROM courses WHERE status = $1 AND deleted_at IS NULL",
		model.ActiveCourseStatus,
	)

	return courses, err
}

func (r *CourseRepository) GetCoursesByUserID(id int) ([]model.Course, error) {
	courses := make([]model.Course, 0)

	err := r.db.Select(
		&courses,
		"SELECT * FROM courses WHERE user_id = $1 AND deleted_at IS NULL",
		id,
	)

	return courses, err
}

func (r *CourseRepository) CreateCourse(course *model.Course) (id int, err error) {
	err = r.db.QueryRow(
		"INSERT INTO courses (uuid, name, description, color, icon, status, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
		course.UUID,
		course.Name,
		course.Description,
		course.Color,
		course.Icon,
		course.Status,
		course.UserID,
		course.CreatedAt,
		course.UpdatedAt,
	).Scan(&id)

	return id, err
}

func (r *CourseRepository) GetCourseByUUID(uuid string) (course model.Course, err error) {
	err = r.db.Get(
		&course,
		"SELECT * FROM courses WHERE uuid = $1 AND deleted_at IS NULL LIMIT 1",
		uuid,
	)

	return course, err
}
