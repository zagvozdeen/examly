package store

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/den4ik117/examly/internal/enum"
	"github.com/guregu/null/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"strings"
	"time"
)

type Question struct {
	ID               int               `json:"id"`
	UUID             string            `json:"uuid"`
	Title            string            `json:"title"`
	Content          null.String       `json:"content"`
	Explanation      null.String       `json:"explanation"`
	ModerationReason null.String       `json:"moderation_reason"`
	Type             enum.QuestionType `json:"type"`
	Status           enum.Status       `json:"status"`
	CourseID         int               `json:"course_id"`
	ModuleID         null.Int          `json:"module_id"`
	CreatedBy        int               `json:"created_by"`
	ModeratedBy      null.Int          `json:"moderated_by"`
	PrevQuestionID   null.Int          `json:"prev_question_id"`
	NextQuestionID   null.Int          `json:"next_question_id"`
	Options          Options           `json:"options"`
	DeletedAt        null.Time         `json:"deleted_at"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
}

type Option struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	IsCorrect bool   `json:"is_correct"`
}

type Options []Option

func (u *Options) Scan(src any) error {
	js, ok := src.(string)
	if !ok {
		return errors.New("source is not a string")
	}
	err := json.Unmarshal([]byte(js), u)
	if err != nil {
		return err
	}
	return nil
}

func (u Options) Value() (driver.Value, error) {
	bytes, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	return string(bytes), nil
}

type QuestionsStore interface {
	Get(ctx context.Context, filter GetQuestionsFilter) ([]Question, error)
	GetByID(ctx context.Context, id int) (Question, error)
	GetByUUID(ctx context.Context, uuid string) (Question, error)
	Create(ctx context.Context, question *Question) error
	GetByCourseID(ctx context.Context, id int) ([]Question, error)
	GetByIDs(ctx context.Context, ids []int) ([]Question, error)
	Update(ctx context.Context, question *Question) error
	UpdateStatus(ctx context.Context, question *Question) error
	Delete(ctx context.Context, question *Question) error
}

type QuestionStore struct {
	conn *pgxpool.Pool
	log  zerolog.Logger
}

type GetQuestionsFilter struct {
	CreatedBy int
	All       bool
}

func (s *QuestionStore) Get(ctx context.Context, filter GetQuestionsFilter) (questions []Question, err error) {
	var sql string
	var params []any

	switch {
	case filter.All:
		sql = "SELECT id, uuid, title, content, explanation, moderation_reason, type, status, course_id, module_id, created_by, moderated_by, prev_question_id, next_question_id, options, deleted_at, created_at, updated_at FROM questions ORDER BY created_at DESC"
		params = []any{}
	case filter.CreatedBy != 0:
		sql = "SELECT id, uuid, title, content, explanation, moderation_reason, type, status, course_id, module_id, created_by, moderated_by, prev_question_id, next_question_id, options, deleted_at, created_at, updated_at FROM questions WHERE created_by = $1 AND deleted_at IS NULL ORDER BY created_at DESC"
		params = []any{filter.CreatedBy}
	default:
		sql = "SELECT id, uuid, title, content, explanation, moderation_reason, type, status, course_id, module_id, created_by, moderated_by, prev_question_id, next_question_id, options, deleted_at, created_at, updated_at FROM questions WHERE status = $1 AND deleted_at IS NULL ORDER BY created_at DESC"
		params = []any{enum.ActiveStatus.String()}
	}

	s.log.Trace().Str("sql", sql).Str("params", fmt.Sprintf("%v", params)).Msg("Query")

	rows, err := s.conn.Query(ctx, sql, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var question Question
		if err := rows.Scan(
			&question.ID,
			&question.UUID,
			&question.Title,
			&question.Content,
			&question.Explanation,
			&question.ModerationReason,
			&question.Type,
			&question.Status,
			&question.CourseID,
			&question.ModuleID,
			&question.CreatedBy,
			&question.ModeratedBy,
			&question.PrevQuestionID,
			&question.NextQuestionID,
			&question.Options,
			&question.DeletedAt,
			&question.CreatedAt,
			&question.UpdatedAt,
		); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}

	return questions, nil
}

func (s *QuestionStore) GetByID(ctx context.Context, id int) (question Question, err error) {
	err = s.conn.QueryRow(
		ctx,
		"SELECT id, options FROM questions WHERE id = $1 AND deleted_at IS NULL",
		id,
	).Scan(
		&question.ID,
		&question.Options,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return question, ErrNotFound
		}
		return question, err
	}

	return question, nil
}

func (s *QuestionStore) GetByUUID(ctx context.Context, uuid string) (question Question, err error) {
	err = s.conn.QueryRow(
		ctx,
		"SELECT id, uuid, title, content, explanation, moderation_reason, type, status, course_id, module_id, created_by, moderated_by, prev_question_id, next_question_id, options, deleted_at, created_at, updated_at FROM questions WHERE uuid = $1 AND deleted_at IS NULL",
		uuid,
	).Scan(
		&question.ID,
		&question.UUID,
		&question.Title,
		&question.Content,
		&question.Explanation,
		&question.ModerationReason,
		&question.Type,
		&question.Status,
		&question.CourseID,
		&question.ModuleID,
		&question.CreatedBy,
		&question.ModeratedBy,
		&question.PrevQuestionID,
		&question.NextQuestionID,
		&question.Options,
		&question.DeletedAt,
		&question.CreatedAt,
		&question.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		err = ErrNotFound
	}

	return question, err
}

func (s *QuestionStore) GetByCourseID(ctx context.Context, id int) (questions []Question, err error) {
	rows, err := s.conn.Query(
		ctx,
		`SELECT id, uuid, title, content, explanation, moderation_reason, type, status, course_id, module_id, created_by, moderated_by, prev_question_id, next_question_id, options, deleted_at, created_at, updated_at
			 FROM questions WHERE course_id = $1 AND deleted_at IS NULL`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var question Question
		if err := rows.Scan(
			&question.ID,
			&question.UUID,
			&question.Title,
			&question.Content,
			&question.Explanation,
			&question.ModerationReason,
			&question.Type,
			&question.Status,
			&question.CourseID,
			&question.ModuleID,
			&question.CreatedBy,
			&question.ModeratedBy,
			&question.PrevQuestionID,
			&question.NextQuestionID,
			&question.Options,
			&question.DeletedAt,
			&question.CreatedAt,
			&question.UpdatedAt,
		); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}

	return questions, nil
}

func (s *QuestionStore) GetByIDs(ctx context.Context, ids []int) (questions []Question, err error) {
	bindings := make([]string, len(ids))
	for i := range ids {
		bindings[i] = fmt.Sprintf("$%d", i+1)
	}
	sql := fmt.Sprintf(
		"SELECT id, uuid, title, content, explanation, moderation_reason, type, status, course_id, module_id, created_by, moderated_by, prev_question_id, next_question_id, options, deleted_at, created_at, updated_at FROM questions WHERE id IN (%s) AND deleted_at IS NULL",
		strings.Join(bindings, ","),
	)
	rows, err := s.conn.Query(ctx, sql, ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var question Question
		if err := rows.Scan(
			&question.ID,
			&question.UUID,
			&question.Title,
			&question.Content,
			&question.Explanation,
			&question.ModerationReason,
			&question.Type,
			&question.Status,
			&question.CourseID,
			&question.ModuleID,
			&question.CreatedBy,
			&question.ModeratedBy,
			&question.PrevQuestionID,
			&question.NextQuestionID,
			&question.Options,
			&question.DeletedAt,
			&question.CreatedAt,
			&question.UpdatedAt,
		); err != nil {
			return nil, err

		}
		questions = append(questions, question)
	}

	return questions, nil
}

func (s *QuestionStore) Create(ctx context.Context, question *Question) error {
	return s.conn.QueryRow(
		ctx,
		`INSERT INTO questions
				(uuid, title, content, explanation, moderation_reason, type, status, course_id, module_id, created_by, moderated_by, prev_question_id, next_question_id, options, created_at, updated_at)
			 VALUES 
				($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
			 RETURNING id`,
		question.UUID,
		question.Title,
		question.Content,
		question.Explanation,
		question.ModerationReason,
		question.Type,
		question.Status,
		question.CourseID,
		question.ModuleID,
		question.CreatedBy,
		question.ModeratedBy,
		question.PrevQuestionID,
		question.NextQuestionID,
		question.Options,
		question.CreatedAt,
		question.UpdatedAt,
	).Scan(&question.ID)
}

func (s *QuestionStore) Update(ctx context.Context, question *Question) error {
	_, err := s.conn.Exec(
		ctx,
		"UPDATE questions SET next_question_id = $1, updated_at = $2 WHERE id = $3 AND deleted_at IS NULL",
		question.NextQuestionID,
		question.UpdatedAt,
		question.ID,
	)
	return err
}

func (s *QuestionStore) UpdateStatus(ctx context.Context, question *Question) error {
	_, err := s.conn.Exec(
		ctx,
		"UPDATE questions SET moderation_reason = $1, status = $2, updated_at = $3, moderated_by = $4 WHERE id = $5 AND deleted_at IS NULL",
		question.ModerationReason,
		question.Status,
		question.UpdatedAt,
		question.ModeratedBy,
		question.ID,
	)
	return err
}

func (s *QuestionStore) Delete(ctx context.Context, question *Question) error {
	_, err := s.conn.Exec(
		ctx,
		"UPDATE questions SET deleted_at = $1, updated_at = $2 WHERE id = $3 AND deleted_at IS NULL",
		question.DeletedAt,
		question.UpdatedAt,
		question.ID,
	)
	return err
}
