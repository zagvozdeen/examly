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
	Get(ctx context.Context, filter GetModulesFilter) ([]Module, error)
	GetByUUID(ctx context.Context, uuid string) (Module, error)
	Create(ctx context.Context, module *Module) error
	Update(ctx context.Context, module *Module) error
	UpdateStatus(ctx context.Context, module *Module) error
	Delete(ctx context.Context, module *Module) error
}

type ModuleStore struct {
	conn *pgxpool.Pool
	log  zerolog.Logger
}

type GetModulesFilter struct {
	CreatedBy   int
	OrCreatedBy int
	All         bool
}

func (s *ModuleStore) Get(ctx context.Context, filter GetModulesFilter) (modules []Module, err error) {
	var sql string
	var params []any

	switch {
	case filter.All:
		sql = "SELECT id, uuid, name, status, course_id, created_by, deleted_at, created_at, updated_at FROM modules"
		params = []any{}
	case filter.CreatedBy != 0:
		sql = "SELECT id, uuid, name, status, course_id, created_by, deleted_at, created_at, updated_at FROM modules WHERE created_by = $1 AND deleted_at IS NULL"
		params = []any{filter.CreatedBy}
	case filter.OrCreatedBy != 0:
		sql = "SELECT id, uuid, name, status, course_id, created_by, deleted_at, created_at, updated_at FROM modules WHERE (created_by = $1 OR status = $2) AND deleted_at IS NULL"
		params = []any{filter.OrCreatedBy, enum.ActiveStatus.String()}
	default:
		sql = "SELECT id, uuid, name, status, course_id, created_by, deleted_at, created_at, updated_at FROM modules WHERE status = $1 AND deleted_at IS NULL"
		params = []any{enum.ActiveStatus.String()}
	}

	s.log.Trace().Str("sql", sql).Str("params", fmt.Sprintf("%v", params)).Msg("Query")

	rows, err := s.conn.Query(ctx, sql, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var module Module
		if err = rows.Scan(
			&module.ID,
			&module.UUID,
			&module.Name,
			&module.Status,
			&module.CourseID,
			&module.CreatedBy,
			&module.DeletedAt,
			&module.CreatedAt,
			&module.UpdatedAt,
		); err != nil {
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
		uuid,
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
