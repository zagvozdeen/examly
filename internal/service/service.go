package service

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/Den4ik117/examly/internal/repository"
)

type Auth interface {
	CreateUser(u RegisterInput) (int, error)
	Login(input LoginInput) (string, error)
	GetGuestToken() (string, error)
	CheckAuth(t string) (*model.User, error)
	UpdateUser(user model.User, u *UpdateUserInput) (*model.User, error)
}

type Courses interface {
	GetCourses() ([]model.Course, error)
	CreateCourse(user *model.User, input *CreateCourseInput) (int, error)
	GetCourseByUUID(uuid string) (model.Course, error)
}

type Modules interface {
	GetModules() ([]model.Module, error)
	CreateModule(input *CreateModuleInput) (int, error)
}

type Questions interface {
	GetQuestions() ([]model.Question, error)
	CreateQuestion(input *CreateQuestionInput) (int, error)
}

type Service struct {
	Auth
	Courses
	Modules
	Questions
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth:      NewAuthService(repos.Auth),
		Courses:   NewCourseService(repos.Courses),
		Modules:   NewModuleService(repos.Modules),
		Questions: NewQuestionService(repos.Questions),
	}
}
