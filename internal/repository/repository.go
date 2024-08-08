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
	UpdateUser(user *model.User) error
}

type Courses interface {
	GetCourses() ([]model.Course, error)
	GetCoursesByUserID(id int) ([]model.Course, error)
	GetAllCourses(id int) ([]model.Course, error)
	GetCoursesByIDs(ids []int) ([]model.Course, error)
	CreateCourse(course *model.Course) (int, error)
	GetCourseByUUID(uuid string) (model.Course, error)
	GetModulesByCourseID(id int) ([]model.Module, error)
	GetQuestionsByCourseID(id int) ([]model.Question, error)
	GetAnswersByIDs(ids []int) ([]model.Answer, error)
	CreateUserQuestions(questions []model.UserQuestion) (map[int]int, error)
	CreateUserCourse(course *model.UserCourse) (int, error)
	CreateUserModules(modules []model.UserModule) error
	CreateUserAnswers(answers []model.UserAnswer) error
}

type Modules interface {
	GetModules() ([]model.Module, error)
	GetModulesByUserID(id int) ([]model.Module, error)
	GetAllModules(id int) ([]model.Module, error)
	CreateModule(module *model.Module) (int, error)
}

type Questions interface {
	GetQuestions() ([]model.Question, error)
	GetQuestionsByUserID(id int) ([]model.Question, error)
	CreateQuestion(question *model.Question) (int, error)
	CreateAnswers(answers []model.Answer) error
	GetFileByID(id int) (model.File, error)
}

type Files interface {
	CreateFile(file *model.File) (id int, err error)
}

type UserCourses interface {
	GetUserCourseByUUID(uuid string) (model.UserCourse, error)
	GetUserQuestionsByCourseID(id int) ([]model.UserQuestion, error)
	GetUserModulesByCourseID(id int) ([]model.UserModule, error)
	GetUserAnswersByQuestionIDs(ids []int) ([]model.UserAnswer, error)
}

type UserQuestionsInterface interface {
	GetUserQuestionByUUID(uuid string) (model.UserQuestion, error)
	GetUserAnswers(id int) ([]model.UserAnswer, error)
	UpdateUserAnswer(answer *model.UserAnswer) error
	UpdateUserQuestion(question *model.UserQuestion) error
	UpdateUserCourse(course *model.UserCourse) error
	CreateUserAnswer(answer *model.UserAnswer) error
}

type Repository struct {
	Auth
	Courses
	Modules
	Questions
	Files
	UserCourses
	UserQuestionsInterface
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:                   NewAuthRepository(db),
		Courses:                NewCourseRepository(db),
		Modules:                NewModuleRepository(db),
		Questions:              NewQuestionRepository(db),
		Files:                  NewFileRepository(db),
		UserCourses:            NewUserCourseRepository(db),
		UserQuestionsInterface: NewUserQuestionRepository(db),
	}
}
