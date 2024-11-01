package store

import (
	"context"
	"fmt"
	"github.com/den4ik117/examly/internal/enum"
	"github.com/guregu/null/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"strings"
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
	Correct        int                  `json:"correct"`
	Incorrect      int                  `json:"incorrect"`
	DeletedAt      null.Time            `json:"deleted_at"`
	CreatedAt      time.Time            `json:"created_at"`
	UpdatedAt      time.Time            `json:"updated_at"`
	Questions      []Question           `json:"questions"`
}

type TestSessionStats struct {
	ID        int
	Correct   int
	Incorrect int
}

type TestSessionsStore interface {
	Get(ctx context.Context, filter GetTestSessionsFilter) ([]TestSession, error)
	GetStats(ctx context.Context, ids []int) ([]TestSessionStats, error)
	GetByID(ctx context.Context, id int) (TestSession, error)
	GetByUUID(ctx context.Context, uuid string) (TestSession, error)
	GetByUserIDAndType(ctx context.Context, id int, t enum.TestSessionType) (TestSession, error)
	Create(ctx context.Context, test *TestSession) error
	Update(ctx context.Context, test *TestSession) error
	GetByCourseID(ctx context.Context, id int) ([]TestSession, error)
}

type TestSessionStore struct {
	conn *pgxpool.Pool
	log  zerolog.Logger
}

type GetTestSessionsFilter struct {
	CourseID int
	UserID   int
}

func (s *TestSessionStore) Get(ctx context.Context, filter GetTestSessionsFilter) (sessions []TestSession, err error) {
	var sql string
	var params []any

	switch {
	case filter.CourseID != 0:
		sql = "SELECT id, uuid, name, type, user_id, course_id, question_ids, last_question_id, deleted_at, created_at, updated_at FROM test_sessions WHERE user_id = $1 AND course_id = $2 AND deleted_at IS NULL"
		params = []any{filter.UserID, filter.CourseID}
	default:
		sql = "SELECT id, uuid, name, type, user_id, course_id, question_ids, last_question_id, deleted_at, created_at, updated_at FROM test_sessions WHERE user_id = $1 AND deleted_at IS NULL"
		params = []any{filter.UserID}
	}

	s.log.Trace().Str("sql", sql).Str("params", fmt.Sprintf("%v", params)).Msg("Query")

	rows, err := s.conn.Query(ctx, sql, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var session TestSession
		err = rows.Scan(
			&session.ID,
			&session.UUID,
			&session.Name,
			&session.Type,
			&session.UserID,
			&session.CourseID,
			&session.QuestionIDs,
			&session.LastQuestionID,
			&session.DeletedAt,
			&session.CreatedAt,
			&session.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}

	return sessions, nil
}

func (s *TestSessionStore) GetStats(ctx context.Context, ids []int) (stats []TestSessionStats, err error) {
	if len(ids) == 0 {
		return
	}

	placeholders := make([]string, len(ids))
	params := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		params[i] = id
	}
	sql := "SELECT test_session_id, COUNT(*) FILTER ( WHERE is_correct IS TRUE ), COUNT(*) FILTER ( WHERE is_correct IS FALSE ) FROM user_answers WHERE test_session_id IN (%s) GROUP BY test_session_id"
	rows, err := s.conn.Query(ctx, fmt.Sprintf(sql, strings.Join(placeholders, ",")), params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var stat TestSessionStats
		err = rows.Scan(
			&stat.ID,
			&stat.Incorrect,
			&stat.Correct,
		)
		if err != nil {
			return nil, err
		}
		stats = append(stats, stat)
	}

	return
}

func (s *TestSessionStore) GetByID(ctx context.Context, id int) (t TestSession, err error) {
	err = s.conn.QueryRow(
		ctx,
		"SELECT id, uuid, name, type, user_id, course_id, question_ids, last_question_id, deleted_at, created_at, updated_at FROM test_sessions WHERE id = $1 AND deleted_at IS NULL",
		id,
	).Scan(
		&t.ID,
		&t.UUID,
		&t.Name,
		&t.Type,
		&t.UserID,
		&t.CourseID,
		&t.QuestionIDs,
		&t.LastQuestionID,
		&t.DeletedAt,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
	return
}

func (s *TestSessionStore) GetByUUID(ctx context.Context, uuid string) (t TestSession, err error) {
	err = s.conn.QueryRow(
		ctx,
		"SELECT id, uuid, name, type, user_id, course_id, question_ids, last_question_id, deleted_at, created_at, updated_at FROM test_sessions WHERE uuid = $1 AND deleted_at IS NULL",
		uuid,
	).Scan(
		&t.ID,
		&t.UUID,
		&t.Name,
		&t.Type,
		&t.UserID,
		&t.CourseID,
		&t.QuestionIDs,
		&t.LastQuestionID,
		&t.DeletedAt,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
	return
}

func (s *TestSessionStore) GetByUserIDAndType(ctx context.Context, id int, t enum.TestSessionType) (ts TestSession, err error) {
	err = s.conn.QueryRow(
		ctx,
		"SELECT id, uuid, name, type, user_id, course_id, question_ids, last_question_id, deleted_at, created_at, updated_at FROM test_sessions WHERE user_id = $1 AND type = $2 AND deleted_at IS NULL",
		id,
		t.String(),
	).Scan(
		&ts.ID,
		&ts.UUID,
		&ts.Name,
		&ts.Type,
		&ts.UserID,
		&ts.CourseID,
		&ts.QuestionIDs,
		&ts.LastQuestionID,
		&ts.DeletedAt,
		&ts.CreatedAt,
		&ts.UpdatedAt,
	)
	return
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
	_, err := s.conn.Exec(
		ctx,
		"UPDATE test_sessions SET name = $1, type = $2, user_id = $3, course_id = $4, question_ids = $5, last_question_id = $6, updated_at = $7 WHERE id = $8",
		test.Name,
		test.Type,
		test.UserID,
		test.CourseID,
		test.QuestionIDs,
		test.LastQuestionID,
		test.UpdatedAt,
		test.ID,
	)
	return err
}

//func (s *TestSessionStore) GetStats(ctx context.Context, userID int) (stats []TestSessionStats, err error) {
//	rows, err := s.conn.Query(
//		ctx,
//		`
//			SELECT c.id                                              AS id,
//				   c.uuid                                            AS uuid,
//				   c.type                                            AS type,
//				   c.created_at                                      AS created_at,
//				   COUNT(q.*) FILTER ( WHERE q.is_correct is true )  AS correct,
//				   COUNT(q.*) FILTER ( WHERE q.is_correct is false ) AS incorrect,
//				   array_length(c.question_ids, 1)                   AS total
//			FROM test_sessions c
//					 JOIN user_answers q on c.id = q.test_session_id
//			WHERE c.user_id = $1 AND c.deleted_at IS NULL
//			GROUP BY c.id, c.uuid, c.type, c.created_at
//			ORDER BY c.id DESC
//`,
//		userID,
//	)
//	if err != nil {
//		return stats, err
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		var stat TestSessionStats
//		err = rows.Scan(
//			&stat.ID,
//			&stat.UUID,
//			&stat.Type,
//			&stat.CreatedAt,
//			&stat.Correct,
//			&stat.Incorrect,
//			&stat.Total,
//		)
//		if err != nil {
//			return stats, err
//		}
//		stats = append(stats, stat)
//	}
//
//	return stats, err
//}

func (s *TestSessionStore) GetByCourseID(ctx context.Context, id int) (ts []TestSession, err error) {
	rows, err := s.conn.Query(
		ctx,
		"SELECT id, uuid, name, type, user_id, course_id, question_ids, last_question_id, deleted_at, created_at, updated_at FROM test_sessions WHERE course_id = $1 AND deleted_at IS NULL",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var t TestSession
		err = rows.Scan(
			&t.ID,
			&t.UUID,
			&t.Name,
			&t.Type,
			&t.UserID,
			&t.CourseID,
			&t.QuestionIDs,
			&t.LastQuestionID,
			&t.DeletedAt,
			&t.CreatedAt,
			&t.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}
	return
}
