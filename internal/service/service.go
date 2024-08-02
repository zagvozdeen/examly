package service

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/Den4ik117/examly/internal/repository"
)

type Courses interface {
	GetCourses() ([]model.Course, error)
}

type Service struct {
	Courses
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Courses: NewCourseService(repos.Courses),
	}
}
