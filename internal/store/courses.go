package store

import (
	"context"
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

type Course struct {
	ID               int           `json:"id"`
	UUID             string        `json:"uuid"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	Color            string        `json:"color"`
	Icon             string        `json:"icon"`
	Status           enum.Status   `json:"status"`
	ModerationReason null.String   `json:"moderation_reason"`
	CreatedBy        int           `json:"created_by"`
	ModeratedBy      null.Int      `json:"moderated_by"`
	DeletedAt        null.Time     `json:"deleted_at"`
	CreatedAt        time.Time     `json:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at"`
	TestSessions     []TestSession `json:"test_sessions"`
}

type CoursesStore interface {
	Get(ctx context.Context, filter GetCoursesFilter) ([]Course, error)
	Create(ctx context.Context, course *Course) error
	GetByUUID(ctx context.Context, uuid string) (Course, error)
	GetByIDs(ctx context.Context, ids []any) ([]Course, error)
	Update(ctx context.Context, course *Course) error
	UpdateStatus(ctx context.Context, course *Course) error
	Delete(ctx context.Context, course *Course) error
}

type CourseStore struct {
	conn *pgxpool.Pool
	log  zerolog.Logger
}

type GetCoursesFilter struct {
	CreatedBy   int
	OrCreatedBy int
	All         bool
}

func (s *CourseStore) Get(ctx context.Context, filter GetCoursesFilter) (courses []Course, err error) {
	var sql string
	var params []any

	switch {
	case filter.All:
		sql = "SELECT id, uuid, name, description, color, icon, status, created_by, deleted_at, created_at, updated_at FROM courses"
		params = []any{}
	case filter.CreatedBy != 0:
		sql = "SELECT id, uuid, name, description, color, icon, status, created_by, deleted_at, created_at, updated_at FROM courses WHERE created_by = $1 AND deleted_at IS NULL"
		params = []any{filter.CreatedBy}
	case filter.OrCreatedBy != 0:
		sql = "SELECT id, uuid, name, description, color, icon, status, created_by, deleted_at, created_at, updated_at FROM courses WHERE (created_by = $1 OR status = $2) AND deleted_at IS NULL"
		params = []any{filter.OrCreatedBy, enum.ActiveStatus.String()}
	default:
		sql = "SELECT id, uuid, name, description, color, icon, status, created_by, deleted_at, created_at, updated_at FROM courses WHERE status = $1 AND deleted_at IS NULL"
		params = []any{enum.ActiveStatus.String()}
	}

	s.log.Trace().Str("sql", sql).Str("params", fmt.Sprintf("%v", params)).Msg("Query")

	rows, err := s.conn.Query(ctx, sql, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var course Course
		err = rows.Scan(
			&course.ID,
			&course.UUID,
			&course.Name,
			&course.Description,
			&course.Color,
			&course.Icon,
			&course.Status,
			&course.CreatedBy,
			&course.DeletedAt,
			&course.CreatedAt,
			&course.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

func (s *CourseStore) Create(ctx context.Context, course *Course) error {
	return s.conn.QueryRow(
		ctx,
		`INSERT INTO courses
				(uuid, name, description, color, icon, status, created_by, created_at, updated_at)
			 VALUES 
				($1, $2, $3, $4, $5, $6, $7, $8, $9)
			 RETURNING id`,
		course.UUID,
		course.Name,
		course.Description,
		course.Color,
		course.Icon,
		course.Status,
		course.CreatedBy,
		course.CreatedAt,
		course.UpdatedAt,
	).Scan(&course.ID)
}

func (s *CourseStore) GetByUUID(ctx context.Context, uuid string) (course Course, err error) {
	err = s.conn.QueryRow(
		ctx,
		"SELECT id, uuid, name, description, color, icon, status, created_by, deleted_at, created_at, updated_at FROM courses WHERE uuid = $1 AND deleted_at IS NULL",
		uuid,
	).Scan(
		&course.ID,
		&course.UUID,
		&course.Name,
		&course.Description,
		&course.Color,
		&course.Icon,
		&course.Status,
		&course.CreatedBy,
		&course.DeletedAt,
		&course.CreatedAt,
		&course.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return course, ErrNotFound
	}

	return course, err
}

func (s *CourseStore) GetByIDs(ctx context.Context, ids []any) (courses []Course, err error) {
	if len(ids) == 0 {
		return courses, nil
	}

	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	sql := fmt.Sprintf("SELECT id, uuid, name, description, color, icon, status, created_by, deleted_at, created_at, updated_at FROM courses WHERE id IN (%s) AND deleted_at IS NULL", strings.Join(placeholders, ","))

	s.log.Trace().Str("sql", sql).Str("params", fmt.Sprintf("%v", ids)).Msg("Query")

	rows, err := s.conn.Query(ctx, sql, ids...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var course Course
		err = rows.Scan(
			&course.ID,
			&course.UUID,
			&course.Name,
			&course.Description,
			&course.Color,
			&course.Icon,
			&course.Status,
			&course.CreatedBy,
			&course.DeletedAt,
			&course.CreatedAt,
			&course.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

func (s *CourseStore) Update(ctx context.Context, course *Course) error {
	_, err := s.conn.Exec(
		ctx,
		"UPDATE courses SET name = $1, description = $2, color = $3, icon = $4, updated_at = $5 WHERE id = $6 AND deleted_at IS NULL",
		course.Name,
		course.Description,
		course.Color,
		course.Icon,
		course.UpdatedAt,
		course.ID,
	)
	return err
}

func (s *CourseStore) UpdateStatus(ctx context.Context, course *Course) error {
	_, err := s.conn.Exec(
		ctx,
		"UPDATE courses SET moderation_reason = $1, status = $2, moderated_by = $3, updated_at = $4 WHERE id = $5 AND deleted_at IS NULL",
		course.ModerationReason,
		course.Status,
		course.ModeratedBy,
		course.UpdatedAt,
		course.ID,
	)
	return err
}

func (s *CourseStore) Delete(ctx context.Context, course *Course) error {
	_, err := s.conn.Exec(
		ctx,
		"UPDATE courses SET updated_at = $1, deleted_at = $2 WHERE id = $3 AND deleted_at IS NULL",
		course.UpdatedAt,
		course.DeletedAt,
		course.ID,
	)
	return err
}
