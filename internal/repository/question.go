package repository

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/jmoiron/sqlx"
)

type QuestionRepository struct {
	db *sqlx.DB
}

func NewQuestionRepository(db *sqlx.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (r *QuestionRepository) GetQuestions() (questions []model.Question, err error) {
	err = r.db.Select(
		&questions,
		"SELECT * FROM questions WHERE status = $1 AND deleted_at IS NULL",
		model.ActiveCourseStatus,
	)

	return questions, err
}

func (r *QuestionRepository) GetQuestionsByUserID(id int) (questions []model.Question, err error) {
	err = r.db.Select(
		&questions,
		"SELECT * FROM questions WHERE user_id = $1 AND deleted_at IS NULL",
		id,
	)

	return questions, err
}

func (r *QuestionRepository) CreateQuestion(question *model.Question) (id int, err error) {
	err = r.db.QueryRow(
		"INSERT INTO questions (uuid, content, explanation, type, status, course_id, module_id, file_id, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id",
		question.UUID,
		question.Content,
		question.Explanation,
		question.Type,
		question.Status,
		question.CourseID,
		question.ModuleID,
		question.FileID,
		question.UserID,
		question.CreatedAt,
		question.UpdatedAt,
	).Scan(&id)

	return id, err
}

func (r *QuestionRepository) CreateAnswers(answers []model.Answer) error {
	_, err := r.db.NamedExec(
		"INSERT INTO answers (content, question_id, is_true, created_at, updated_at) VALUES (:content, :question_id, :is_true, :created_at, :updated_at)",
		answers,
	)

	return err
}
