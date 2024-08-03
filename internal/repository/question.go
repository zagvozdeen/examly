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
		"SELECT * FROM questions WHERE deleted_at IS NULL",
	)

	return questions, err
}

func (r *QuestionRepository) CreateQuestion(question *model.Question) (id int, err error) {
	err = r.db.QueryRow(
		"INSERT INTO questions (content, type, course_id, module_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		question.Content,
		question.Type,
		question.CourseID,
		question.ModuleID,
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
