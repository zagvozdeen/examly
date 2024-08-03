package service

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/Den4ik117/examly/internal/repository"
	"time"
)

type ModuleService struct {
	repo repository.Modules
}

type CreateModuleInput struct {
	Name     string `json:"name" validate:"required,max=255"`
	CourseID int    `json:"course_id" validate:"required"`
}

func NewModuleService(repo repository.Modules) *ModuleService {
	return &ModuleService{repo: repo}
}

func (s *ModuleService) GetModules() ([]model.Module, error) {
	return s.repo.GetModules()
}

func (s *ModuleService) CreateModule(input *CreateModuleInput) (int, error) {
	course := &model.Module{
		Name:      input.Name,
		CourseID:  input.CourseID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.repo.CreateModule(course)
}
