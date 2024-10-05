package store

import (
	"context"
	"github.com/den4ik117/examly/internal/enum"
	"github.com/guregu/null/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Module struct {
	ID        int                `json:"id"`
	UUID      string             `json:"uuid"`
	Name      string             `json:"name"`
	Status    enum.Status        `json:"status"`
	CourseID  int                `json:"course_id"`
	CreatedBy int                `json:"created_by"`
	DeletedAt null.Time          `json:"deleted_at"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Course    null.Value[Course] `json:"course"`
}

type ModulesStore interface {
	Get(ctx context.Context) ([]Module, error)
	Create(ctx context.Context, course *Module) error
}

type ModuleStore struct {
	conn *pgxpool.Pool
}

func (s *ModuleStore) Get(ctx context.Context) (modules []Module, err error) {
	rows, err := s.conn.Query(
		ctx,
		"SELECT id, uuid, name, status, course_id, created_by, deleted_at, created_at, updated_at FROM modules WHERE deleted_at IS NULL",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var module Module
		err = rows.Scan(
			&module.ID,
			&module.UUID,
			&module.Name,
			&module.Status,
			&module.CourseID,
			&module.CreatedBy,
			&module.DeletedAt,
			&module.CreatedAt,
			&module.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		modules = append(modules, module)
	}

	return modules, nil
}

func (s *ModuleStore) Create(ctx context.Context, module *Module) error {
	return s.conn.QueryRow(
		ctx,
		`INSERT INTO modules
				(uuid, name, status, course_id, created_by, created_at, updated_at)
			 VALUES 
				($1, $2, $3, $4, $5, $6, $7)
			 RETURNING id`,
		module.UUID,
		module.Name,
		module.Status,
		module.CourseID,
		module.CreatedBy,
		module.CreatedAt,
		module.UpdatedAt,
	).Scan(&module.ID)
}
