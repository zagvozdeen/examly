package service

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/Den4ik117/examly/internal/repository"
	"github.com/google/uuid"
	"time"
)

type CourseService struct {
	repo repository.Courses
}

type CreateCourseInput struct {
	Name string `json:"name" validate:"required,max=255"`
}

func NewCourseService(repo repository.Courses) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) GetCourses() ([]model.Course, error) {
	return s.repo.GetCourses()
}

func (s *CourseService) CreateCourse(user *model.User, input *CreateCourseInput) (int, error) {
	courseUUID, err := uuid.NewV7()
	if err != nil {
		return 0, err
	}

	course := &model.Course{
		UUID:      courseUUID.String(),
		Name:      input.Name,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.repo.CreateCourse(course)
}

func (s *CourseService) GetCourseByUUID(uuid string) (model.Course, error) {
	return s.repo.GetCourseByUUID(uuid)
}
