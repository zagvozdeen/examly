package service

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/Den4ik117/examly/internal/repository"
	"github.com/google/uuid"
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

func (s *ModuleService) GetModulesByUserID(id int) ([]model.Module, error) {
	return s.repo.GetModulesByUserID(id)
}

func (s *ModuleService) GetAllModules(id int) ([]model.Module, error) {
	return s.repo.GetAllModules(id)
}

func (s *ModuleService) CreateModule(user *model.User, input *CreateModuleInput) (int, error) {
	uid, err := uuid.NewV7()
	if err != nil {
		return 0, err
	}

	course := &model.Module{
		UUID:      uid.String(),
		Name:      input.Name,
		Status:    model.NewCourseStatus,
		CourseID:  input.CourseID,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.repo.CreateModule(course)
}
