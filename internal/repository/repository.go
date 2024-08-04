package repository

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	CreateUser(user model.User) (int, error)
	GetUserByEmail(email string) (model.User, error)
	IsExistsUserByEmail(email string) (bool, error)
	GetUserByID(id int) (model.User, error)
}

type Courses interface {
	GetCourses() ([]model.Course, error)
	CreateCourse(course *model.Course) (int, error)
	GetCourseByUUID(uuid string) (model.Course, error)
}

type Modules interface {
	GetModules() ([]model.Module, error)
	CreateModule(module *model.Module) (int, error)
}

type Questions interface {
	GetQuestions() ([]model.Question, error)
	CreateQuestion(question *model.Question) (int, error)
	CreateAnswers(answers []model.Answer) error
}

type Repository struct {
	Auth
	Courses
	Modules
	Questions
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:      NewAuthRepository(db),
		Courses:   NewCourseRepository(db),
		Modules:   NewModuleRepository(db),
		Questions: NewQuestionRepository(db),
	}
}
