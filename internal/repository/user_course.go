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

func (r *UserCourseRepository) GetUserCourseByID(id int) (course model.UserCourse, err error) {
	err = r.db.Get(
		&course,
		"SELECT * FROM user_courses WHERE id = $1 AND deleted_at IS NULL LIMIT 1",
		id,
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

func (r *UserCourseRepository) GetUserCourseByTypeAndUserID(t string, id int) (course model.UserCourse, err error) {
	err = r.db.Get(
		&course,
		"SELECT * FROM user_courses WHERE type = $1 AND user_id = $2 AND deleted_at IS NULL LIMIT 1",
		t,
		id,
	)

	return course, err
}

func (r *UserCourseRepository) CreateUserQuestion(question *model.UserQuestion) error {
	return r.db.QueryRowx(
		"INSERT INTO user_questions (uuid, content, explanation, type, is_true, sort, course_id, question_id, module_id, file_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id",
		question.UUID,
		question.Content,
		question.Explanation,
		question.Type,
		question.IsTrue,
		question.Sort,
		question.CourseID,
		question.QuestionID,
		question.ModuleID,
		question.FileID,
		question.CreatedAt,
		question.UpdatedAt,
	).Scan(&question.ID)
}
