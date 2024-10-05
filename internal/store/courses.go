package store

import (
	"context"
	"errors"
	"github.com/den4ik117/examly/internal/enum"
	"github.com/den4ik117/examly/internal/util"
	"github.com/guregu/null/v5"
	"github.com/jackc/pgx/v5"
	"time"
)

type Course struct {
	ID          int         `json:"id"`
	UUID        string      `json:"uuid"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Color       string      `json:"color"`
	Icon        string      `json:"icon"`
	Status      enum.Status `json:"status"`
	CreatedBy   int         `json:"created_by"`
	DeletedAt   null.Time   `json:"deleted_at"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type CoursesStore interface {
	Get(ctx context.Context, filter GetCoursesFilter) ([]Course, error)
	Create(ctx context.Context, course *Course) error
	GetByUUID(ctx context.Context, uuid string) (Course, error)
}

type CourseStore struct {
	conn *pgx.Conn
}

type GetCoursesFilter struct {
	Trashed  bool
	Statuses []any
	UserID   int
}

func (s *CourseStore) Get(ctx context.Context, filter GetCoursesFilter) (courses []Course, err error) {
	b := util.NewQueryBuilder("SELECT id, uuid, name, description, color, icon, status, created_by, deleted_at, created_at, updated_at FROM courses")
	if !filter.Trashed {
		b.WhereNull("deleted_at")
	}
	if filter.UserID != 0 && filter.Statuses != nil {
		b.WhereFunc(func(b *util.QueryBuilder) {
			b.WhereIn("status", filter.Statuses).OrWhere("user_id", "=", filter.UserID)
		})
	} else if filter.UserID != 0 {
		b.Where("user_id", "=", filter.UserID)
	}

	rows, err := s.conn.Query(ctx, b.String(), b.Params()...)
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
