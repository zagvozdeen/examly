package service

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/Den4ik117/examly/internal/repository"
	"mime/multipart"
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
	GetCoursesByUserID(id int) ([]model.Course, error)
	GetAllCourses(id int) ([]model.Course, error)
	GetModuleCourses(modules []model.Module) ([]model.Course, error)
	CreateCourse(user *model.User, input *CreateCourseInput) (int, error)
	GetCourseByUUID(uuid string) (model.Course, error)
	CreateUserCourse(input *CreateUserCourseInput) (string, error)
	GetQuestionsAnswers(questions []model.Question) ([]model.Answer, error)
}

type Modules interface {
	GetModules() ([]model.Module, error)
	GetModulesByUserID(id int) ([]model.Module, error)
	GetAllModules(id int) ([]model.Module, error)
	CreateModule(user *model.User, input *CreateModuleInput) (int, error)
}

type Questions interface {
	GetQuestions() ([]model.Question, error)
	GetQuestionsByUserID(id int) ([]model.Question, error)
	CreateQuestion(user *model.User, input *CreateQuestionInput) (int, error)
	ImportQuestions(input *ImportQuestionsInput) error
}

type Files interface {
	UploadFile(user *model.User, file multipart.File, header *multipart.FileHeader) (*model.File, error)
}

type UserCourses interface {
	GetUserCourseByUUID(uuid string) (model.UserCourse, error)
}

type UserQuestions interface {
	CheckAnswer(input *CheckAnswerInput) (*model.UserQuestion, error)
}

type Service struct {
	Auth
	Courses
	Modules
	Questions
	Files
	UserCourses
	UserQuestions
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth:          NewAuthService(repos.Auth),
		Courses:       NewCourseService(repos.Courses),
		Modules:       NewModuleService(repos.Modules),
		Questions:     NewQuestionService(repos.Questions),
		Files:         NewFileService(repos.Files),
		UserCourses:   NewUserCourseService(repos.UserCourses),
		UserQuestions: NewUserQuestionService(repos.UserQuestionsInterface),
	}
}
