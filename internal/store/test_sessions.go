package store

import (
	"context"
	"github.com/den4ik117/examly/internal/enum"
	"github.com/guregu/null/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type TestSession struct {
	ID             int                  `json:"id"`
	UUID           string               `json:"uuid"`
	Name           string               `json:"name"`
	Type           enum.TestSessionType `json:"type"`
	UserID         int                  `json:"user_id"`
	CourseID       null.Int             `json:"course_id"`
	QuestionIDs    []int                `json:"question_ids"`
	LastQuestionID null.Int             `json:"last_question_id"`
	DeletedAt      null.Time            `json:"deleted_at"`
	CreatedAt      time.Time            `json:"created_at"`
	UpdatedAt      time.Time            `json:"updated_at"`
}

type TestSessionStats struct {
	ID            int       `json:"id"`
	UUID          string    `json:"uuid"`
	Type          string    `json:"type"`
	CreatedAt     time.Time `json:"-"`
	CreatedAtUnix int64     `json:"created_at"`
	Correct       int       `json:"correct"`
	Incorrect     int       `json:"incorrect"`
	Total         int       `json:"total"`
}

type TestSessionsStore interface {
	GetByID(ctx context.Context, id int) (TestSession, error)
	GetByUUID(ctx context.Context, uuid string) (TestSession, error)
	GetByUserIDAndType(ctx context.Context, id int, t enum.TestSessionType) (TestSession, error)
	Create(ctx context.Context, test *TestSession) error
	Update(ctx context.Context, test *TestSession) error
	GetStats(ctx context.Context, userID int) ([]TestSessionStats, error)
}

type TestSessionStore struct {
	conn *pgxpool.Pool
}

func (s *TestSessionStore) GetByID(ctx context.Context, id int) (TestSession, error) {
	return TestSession{}, nil
}

func (s *TestSessionStore) GetByUUID(ctx context.Context, uuid string) (TestSession, error) {
	return TestSession{}, nil
}

func (s *TestSessionStore) GetByUserIDAndType(ctx context.Context, id int, t enum.TestSessionType) (TestSession, error) {
	return TestSession{}, nil
}

func (s *TestSessionStore) Create(ctx context.Context, test *TestSession) error {
	return s.conn.QueryRow(
		ctx,
		"INSERT INTO test_sessions (uuid, name, type, user_id, course_id, question_ids, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		test.UUID,
		test.Name,
		test.Type,
		test.UserID,
		test.CourseID,
		test.QuestionIDs,
		test.CreatedAt,
		test.UpdatedAt,
	).Scan(&test.ID)
}

func (s *TestSessionStore) Update(ctx context.Context, test *TestSession) error {
	return nil
}

func (s *TestSessionStore) GetStats(ctx context.Context, userID int) (stats []TestSessionStats, err error) {
	rows, err := s.conn.Query(
		ctx,
		`
			SELECT c.id                                              AS id,
				   c.uuid                                            AS uuid,
				   c.type                                            AS type,
				   c.created_at                                      AS created_at,
				   COUNT(q.*) FILTER ( WHERE q.is_correct is true )  AS correct,
				   COUNT(q.*) FILTER ( WHERE q.is_correct is false ) AS incorrect,
				   array_length(c.question_ids, 1)                   AS total
			FROM test_sessions c
					 JOIN user_answers q on c.id = q.test_session_id
			WHERE c.user_id = $1 AND c.deleted_at IS NULL
			GROUP BY c.id, c.uuid, c.type, c.created_at
			ORDER BY c.id DESC
`,
		userID,
	)
	if err != nil {
		return stats, err
	}
	defer rows.Close()

	for rows.Next() {
		var stat TestSessionStats
		err = rows.Scan(
			&stat.ID,
			&stat.UUID,
			&stat.Type,
			&stat.CreatedAt,
			&stat.Correct,
			&stat.Incorrect,
			&stat.Total,
		)
		if err != nil {
			return stats, err
		}
		stats = append(stats, stat)
	}

	return stats, err
}
