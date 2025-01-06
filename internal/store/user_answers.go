package store

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type UserAnswer struct {
	ID            int            `json:"id"`
	TestSessionID int            `json:"test_session_id"`
	QuestionID    int            `json:"question_id"`
	AnswerData    map[string]any `json:"answer_data"`
	IsCorrect     bool           `json:"is_correct"`
	AnsweredAt    time.Time      `json:"answered_at"`
}

type UserAnswersStore interface {
	GetByTestSessionID(ctx context.Context, id int) ([]UserAnswer, error)
	Create(ctx context.Context, answer *UserAnswer) error
}

type UserAnswerStore struct {
	conn *pgxpool.Pool
}

func (s *UserAnswerStore) GetByTestSessionID(ctx context.Context, id int) (answers []UserAnswer, err error) {
	rows, err := s.conn.Query(
		ctx,
		"SELECT id, test_session_id, question_id, answer_data, is_correct, answered_at FROM user_answers WHERE test_session_id = $1 ORDER BY answered_at",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var answer UserAnswer
		err = rows.Scan(
			&answer.ID,
			&answer.TestSessionID,
			&answer.QuestionID,
			&answer.AnswerData,
			&answer.IsCorrect,
			&answer.AnsweredAt,
		)
		if err != nil {
			return nil, err
		}
		answers = append(answers, answer)
	}
	return
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
