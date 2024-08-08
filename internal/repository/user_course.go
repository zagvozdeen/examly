package repository

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/jmoiron/sqlx"
)

type UserCourseRepository struct {
	db *sqlx.DB
}

func NewUserCourseRepository(db *sqlx.DB) *UserCourseRepository {
	return &UserCourseRepository{db: db}
}

func (r *UserCourseRepository) GetUserCourseByUUID(uuid string) (course model.UserCourse, err error) {
	err = r.db.Get(
		&course,
		"SELECT * FROM user_courses WHERE uuid = $1 AND deleted_at IS NULL LIMIT 1",
		uuid,
	)

	return course, err
}

func (r *UserCourseRepository) GetUserQuestionsByCourseID(id int) (questions []model.UserQuestion, err error) {
	err = r.db.Select(
		&questions,
		"SELECT * FROM user_questions WHERE course_id = $1 AND deleted_at IS NULL ORDER BY sort",
		id,
	)

	return questions, err
}

func (r *UserCourseRepository) GetUserModulesByCourseID(id int) (modules []model.UserModule, err error) {
	err = r.db.Select(
		&modules,
		"SELECT * FROM user_modules WHERE course_id = $1 AND deleted_at IS NULL",
		id,
	)

	return modules, err
}

func (r *UserCourseRepository) GetUserAnswersByQuestionIDs(ids []int) ([]model.UserAnswer, error) {
	answers := make([]model.UserAnswer, 0)

	query, args, err := sqlx.In("SELECT * FROM user_answers WHERE question_id IN (?) AND deleted_at IS NULL ORDER BY sort", ids)
	query = r.db.Rebind(query)
	err = r.db.Select(&answers, query, args...)

	return answers, err
}
