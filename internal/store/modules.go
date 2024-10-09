package store

import (
	"context"
	"errors"
	"github.com/den4ik117/examly/internal/enum"
	"github.com/guregu/null/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Module struct {
	ID               int                `json:"id"`
	UUID             string             `json:"uuid"`
	Name             string             `json:"name"`
	Status           enum.Status        `json:"status"`
	ModerationReason null.String        `json:"moderation_reason"`
	CourseID         int                `json:"course_id"`
	CreatedBy        int                `json:"created_by"`
	ModeratedBy      null.Int           `json:"moderated_by"`
	DeletedAt        null.Time          `json:"deleted_at"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
	Course           null.Value[Course] `json:"course"`
}

type ModulesStore interface {
	Get(ctx context.Context) ([]Module, error)
	GetByUUID(ctx context.Context, uuid string) (Module, error)
	GetByCreatedBy(ctx context.Context, id int) ([]Module, error)
	Create(ctx context.Context, module *Module) error
	Update(ctx context.Context, module *Module) error
	UpdateStatus(ctx context.Context, module *Module) error
	Delete(ctx context.Context, module *Module) error
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

func (s *ModuleStore) GetByUUID(ctx context.Context, uuid string) (module Module, err error) {
	err = s.conn.QueryRow(
		ctx,
		"SELECT id, uuid, name, status, course_id, created_by, deleted_at, created_at, updated_at FROM modules WHERE uuid = $1 AND deleted_at IS NULL",
	).Scan(
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

	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		err = ErrNotFound
	}

	return
}

func (s *ModuleStore) GetByCreatedBy(ctx context.Context, id int) (modules []Module, err error) {
	rows, err := s.conn.Query(
		ctx,
		"SELECT id, uuid, name, status, course_id, created_by, deleted_at, created_at, updated_at FROM modules WHERE created_by = $1 AND deleted_at IS NULL",
		id,
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

func (s *ModuleStore) Update(ctx context.Context, module *Module) error {
	_, err := s.conn.Exec(
		ctx,
		`UPDATE modules SET name = $1, updated_at = $2 WHERE id = $3 AND deleted_at IS NULL`,
		module.Name,
		module.UpdatedAt,
		module.ID,
	)
	return err
}

func (s *ModuleStore) UpdateStatus(ctx context.Context, module *Module) error {
	_, err := s.conn.Exec(
		ctx,
		`UPDATE modules SET moderation_reason = $1, status = $2, moderated_by = $3, updated_at = $4 WHERE id = $5 AND deleted_at IS NULL`,
		module.ModerationReason,
		module.Status,
		module.ModeratedBy,
		module.UpdatedAt,
		module.ID,
	)
	return err
}

func (s *ModuleStore) Delete(ctx context.Context, module *Module) error {
	_, err := s.conn.Exec(
		ctx,
		`UPDATE modules SET deleted_at = $1, updated_at = $2 WHERE id = $3 AND deleted_at IS NULL`,
		module.DeletedAt,
		module.UpdatedAt,
		module.ID,
	)
	return err
}
