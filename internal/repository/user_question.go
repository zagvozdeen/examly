package repository

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/jmoiron/sqlx"
)

type UserQuestionRepository struct {
	db *sqlx.DB
}

func NewUserQuestionRepository(db *sqlx.DB) *UserQuestionRepository {
	return &UserQuestionRepository{db: db}
}

func (r *UserQuestionRepository) GetUserQuestionByUUID(uuid string) (question model.UserQuestion, err error) {
	err = r.db.Get(
		&question,
		"SELECT * FROM user_questions WHERE uuid = $1 AND deleted_at IS NULL LIMIT 1",
		uuid,
	)

	return question, err
}

func (r *UserQuestionRepository) GetUserAnswers(id int) (answers []model.UserAnswer, err error) {
	err = r.db.Select(
		&answers,
		"SELECT * FROM user_answers WHERE question_id = $1 AND deleted_at IS NULL ORDER BY sort",
		id,
	)

	return answers, err
}

func (r *UserQuestionRepository) UpdateUserAnswer(answer *model.UserAnswer) (err error) {
	_, err = r.db.Exec(
		"UPDATE user_answers SET is_chosen = $1, updated_at = $2 WHERE id = $3 AND deleted_at IS NULL",
		answer.IsChosen,
		answer.UpdatedAt,
		answer.ID,
	)

	return err
}

func (r *UserQuestionRepository) UpdateUserQuestion(question *model.UserQuestion) (err error) {
	_, err = r.db.Exec(
		"UPDATE user_questions SET is_true = $1, updated_at = $2 WHERE id = $3 AND deleted_at IS NULL",
		question.IsTrue,
		question.UpdatedAt,
		question.ID,
	)

	return err
}

func (r *UserQuestionRepository) UpdateUserCourse(course *model.UserCourse) (err error) {
	_, err = r.db.Exec(
		"UPDATE user_courses SET last_question_id = $1, updated_at = $2 WHERE id = $3 AND deleted_at IS NULL",
		course.LastQuestionID,
		course.UpdatedAt,
		course.ID,
	)

	return err
}

func (r *UserQuestionRepository) CreateUserAnswer(answer *model.UserAnswer) error {
	_, err := r.db.NamedQuery(
		"INSERT INTO user_answers (content, question_id, is_true, is_chosen, sort, created_at, updated_at) VALUES (:content, :question_id, :is_true, :is_chosen, :sort, :created_at, :updated_at)",
		answer,
	)

	return err
}
