package service

import (
	"fmt"
	"github.com/Den4ik117/examly/internal/model"
	"github.com/Den4ik117/examly/internal/repository"
	"github.com/Den4ik117/examly/internal/util"
	"github.com/google/uuid"
	"slices"
	"time"
)

type CourseService struct {
	repo repository.Courses
}

type CreateCourseInput struct {
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description" validate:"required"`
	Color       string `json:"color" validate:"required,max=255"`
	Icon        string `json:"icon" validate:"required,max=255"`
}

type CreateUserCourseInput struct {
	CourseUUID string
	ModuleID   int
	UserID     int
	Type       string
}

func NewCourseService(repo repository.Courses) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) GetCourses() ([]model.Course, error) {
	return s.repo.GetCourses()
}

func (s *CourseService) GetCoursesByUserID(id int) ([]model.Course, error) {
	return s.repo.GetCoursesByUserID(id)
}

func (s *CourseService) GetAllCourses(id int) ([]model.Course, error) {
	return s.repo.GetAllCourses(id)
}

func (s *CourseService) GetModuleCourses(modules []model.Module) ([]model.Course, error) {
	ids := make([]int, len(modules))
	for i, module := range modules {
		ids[i] = module.CourseID
	}
	ids = util.UniqueIntSlice(ids)

	return s.repo.GetCoursesByIDs(ids)
}

func (s *CourseService) CreateCourse(user *model.User, input *CreateCourseInput) (int, error) {
	courseUUID, err := uuid.NewV7()
	if err != nil {
		return 0, err
	}

	course := &model.Course{
		UUID:        courseUUID.String(),
		Name:        input.Name,
		Description: input.Description,
		UserID:      user.ID,
		Color:       input.Color,
		Icon:        input.Icon,
		Status:      model.NewCourseStatus,
	}
	course.FillTime()

	return s.repo.CreateCourse(course)
}

func (s *CourseService) CreateUserCourse(input *CreateUserCourseInput) (string, error) {
	if !slices.Contains(model.AllUserCourseTypes, input.Type) {
		return "", fmt.Errorf("invalid user course type")
	}

	course, err := s.repo.GetCourseByUUID(input.CourseUUID)
	if err != nil {
		return "", err
	}

	modules, err := s.repo.GetModulesByCourseID(course.ID)
	if err != nil {
		return "", err
	}

	questions, err := s.repo.GetQuestionsByCourseID(course.ID)
	if err != nil {
		return "", err
	}

	answers, err := s.GetQuestionsAnswers(questions)
	if err != nil {
		return "", err
	}

	uid, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	userCourse := &model.UserCourse{
		Model: model.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UUID:     uid.String(),
		Name:     course.Name,
		Type:     input.Type,
		UserID:   input.UserID,
		CourseID: course.ID,
	}

	id, err := s.repo.CreateUserCourse(userCourse)
	if err != nil {
		return "", err
	}
	userCourse.ID = id

	userModules := make([]model.UserModule, len(modules))
	for i, module := range modules {
		userModules[i] = model.UserModule{
			Model: model.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:     module.Name,
			CourseID: userCourse.ID,
		}
	}

	err = s.repo.CreateUserModules(userModules)
	if err != nil {
		return "", err
	}

	userQuestions := make([]model.UserQuestion, len(questions))
	order := util.RandomIntSlice(len(questions))
	for i, question := range questions {
		quuid, err := uuid.NewV7()
		if err != nil {
			return "", err
		}
		userQuestions[i] = model.UserQuestion{
			Model: model.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			UUID:        quuid.String(),
			Content:     question.Content,
			Explanation: question.Explanation,
			Type:        question.Type,
			Sort:        order[i],
			CourseID:    userCourse.ID,
			QuestionID:  question.ID,
			ModuleID:    question.ModuleID,
			FileID:      question.FileID,
		}
	}

	qids, err := s.repo.CreateUserQuestions(userQuestions)
	if err != nil {
		return "", err
	}

	userAnswers := make([]model.UserAnswer, len(answers))
	order = util.RandomIntSlice(len(answers))
	for i, answer := range answers {
		userAnswers[i] = model.UserAnswer{
			Model: model.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Content:    answer.Content,
			QuestionID: qids[answer.QuestionID],
			IsTrue:     answer.IsTrue,
			IsChosen:   false,
			Sort:       order[i],
		}
	}

	err = s.repo.CreateUserAnswers(userAnswers)
	if err != nil {
		return "", err
	}

	return uid.String(), nil
}

func (s *CourseService) GetQuestionsAnswers(questions []model.Question) ([]model.Answer, error) {
	ids := make([]int, len(questions))
	for i, question := range questions {
		ids[i] = question.ID
	}

	return s.repo.GetAnswersByIDs(ids)
}

func (s *CourseService) GetCourseByUUID(uuid string) (model.Course, error) {
	return s.repo.GetCourseByUUID(uuid)
}
