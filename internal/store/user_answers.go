package store

import (
	"context"
	"github.com/jackc/pgx/v5"
	"time"
)

type UserAnswer struct {
	ID            int       `json:"id"`
	TestSessionID int       `json:"test_session_id"`
	QuestionID    int       `json:"question_id"`
	AnswerData    string    `json:"answer_data"`
	IsCorrect     bool      `json:"is_correct"`
	AnsweredAt    time.Time `json:"answered_at"`
}

type UserAnswersStore interface {
	Create(ctx context.Context, answer *UserAnswer) error
}

type UserAnswerStore struct {
	conn *pgx.Conn
}

func (s *UserAnswerStore) Create(ctx context.Context, answer *UserAnswer) error {
	return s.conn.QueryRow(
		ctx,
		`INSERT INTO user_answers
				(test_session_id, question_id, answer_data, is_correct, answered_at)
			 VALUES
				($1, $2, $3, $4, $5)
			 RETURNING id`,
		answer.TestSessionID,
		answer.QuestionID,
		answer.AnswerData,
		answer.IsCorrect,
		answer.AnsweredAt,
	).Scan(&answer.ID)
}
