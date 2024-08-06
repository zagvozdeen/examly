package service

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/Den4ik117/examly/internal/repository"
	"github.com/Den4ik117/examly/internal/util"
	"github.com/google/uuid"
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

func (s *CourseService) GetCourseByUUID(uuid string) (model.Course, error) {
	return s.repo.GetCourseByUUID(uuid)
}
