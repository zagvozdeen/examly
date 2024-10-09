package store

import (
	"context"
	"errors"
	"github.com/den4ik117/examly/internal/enum"
	"github.com/den4ik117/examly/internal/util"
	"github.com/guregu/null/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Course struct {
	ID               int         `json:"id"`
	UUID             string      `json:"uuid"`
	Name             string      `json:"name"`
	Description      string      `json:"description"`
	Color            string      `json:"color"`
	Icon             string      `json:"icon"`
	Status           enum.Status `json:"status"`
	ModerationReason null.String `json:"moderation_reason"`
	CreatedBy        int         `json:"created_by"`
	ModeratedBy      null.Int    `json:"moderated_by"`
	DeletedAt        null.Time   `json:"deleted_at"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
}

type CoursesStore interface {
	Get(ctx context.Context, filter GetCoursesFilter) ([]Course, error)
	Create(ctx context.Context, course *Course) error
	GetByUUID(ctx context.Context, uuid string) (Course, error)
	Update(ctx context.Context, course *Course) error
	UpdateStatus(ctx context.Context, course *Course) error
	Delete(ctx context.Context, course *Course) error
}

type CourseStore struct {
	conn *pgxpool.Pool
}

type GetCoursesFilter struct {
	Trashed   bool
	Statuses  []any
	CreatedBy int
}

func (s *CourseStore) Get(ctx context.Context, filter GetCoursesFilter) (courses []Course, err error) {
	b := util.NewQueryBuilder("SELECT id, uuid, name, description, color, icon, status, created_by, deleted_at, created_at, updated_at FROM courses")
	if !filter.Trashed {
		b.WhereNull("deleted_at")
	}
	if filter.CreatedBy != 0 && filter.Statuses != nil {
		b.WhereFunc(func(b *util.QueryBuilder) {
			b.WhereIn("status", filter.Statuses).OrWhere("user_id", "=", filter.CreatedBy)
		})
	} else if filter.CreatedBy != 0 {
		b.Where("created_by", "=", filter.CreatedBy)
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
