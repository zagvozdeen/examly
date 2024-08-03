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

func (r *ModuleRepository) GetModules() (modules []model.Module, err error) {
	err = r.db.Select(
		&modules,
		"SELECT * FROM modules WHERE deleted_at IS NULL",
	)

	return modules, err
}

func (r *ModuleRepository) CreateModule(module *model.Module) (id int, err error) {
	err = r.db.QueryRow(
		"INSERT INTO modules (name, course_id, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id",
		module.Name,
		module.CourseID,
		module.CreatedAt,
		module.UpdatedAt,
	).Scan(&id)

	return id, err
}
