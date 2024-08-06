package repository

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/jmoiron/sqlx"
)

type ModuleRepository struct {
	db *sqlx.DB
}

func NewModuleRepository(db *sqlx.DB) *ModuleRepository {
	return &ModuleRepository{db: db}
}

func (r *ModuleRepository) GetModules() ([]model.Module, error) {
	modules := make([]model.Module, 0)

	err := r.db.Select(
		&modules,
		"SELECT * FROM modules WHERE status = $1 AND deleted_at IS NULL",
		model.ActiveCourseStatus,
	)

	return modules, err
}

func (r *ModuleRepository) GetModulesByUserID(id int) ([]model.Module, error) {
	modules := make([]model.Module, 0)

	err := r.db.Select(
		&modules,
		"SELECT * FROM modules WHERE user_id = $1 AND deleted_at IS NULL",
		id,
	)

	return modules, err
}

func (r *ModuleRepository) GetAllModules(id int) ([]model.Module, error) {
	modules := make([]model.Module, 0)

	err := r.db.Select(
		&modules,
		"SELECT * FROM modules WHERE (user_id = $1 OR status = $2) AND deleted_at IS NULL",
		id,
		model.ActiveCourseStatus,
	)

	return modules, err
}

func (r *ModuleRepository) CreateModule(module *model.Module) (id int, err error) {
	err = r.db.QueryRow(
		"INSERT INTO modules (uuid, name, status, course_id, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		module.UUID,
		module.Name,
		module.Status,
		module.CourseID,
		module.UserID,
		module.CreatedAt,
		module.UpdatedAt,
	).Scan(&id)

	return id, err
}
