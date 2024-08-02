package service

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/Den4ik117/examly/internal/repository"
)

type CourseService struct {
	repo repository.Courses
}

func NewCourseService(repo repository.Courses) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) GetCourses() ([]model.Course, error) {
	return s.repo.GetCourses()
}
