package repository

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	CreateUser(user model.User) (int, error)
	GetUserByEmail(email string) (model.User, error)
	IsExistsUserByEmail(email string) (bool, error)
	GetUserByID(id int) (user model.User, err error)
}

type Courses interface {
	GetCourses() ([]model.Course, error)
}

type Repository struct {
	Auth
	Courses
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:    NewAuthRepository(db),
		Courses: NewCourseRepository(db),
	}
}
