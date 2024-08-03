package service

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/Den4ik117/examly/internal/repository"
	"github.com/guregu/null/v5"
	"time"
)

type QuestionService struct {
	repo repository.Questions
}

type Answers struct {
	Content string `json:"content" validate:"required"`
	IsTrue  bool   `json:"is_true" validate:""`
}

type CreateQuestionInput struct {
	Content  string    `json:"content" validate:"required"`
	Type     string    `json:"type" validate:"required"`
	CourseID int       `json:"course_id" validate:"required"`
	ModuleID int       `json:"module_id" validate:"required"`
	Answers  []Answers `json:"answers" validate:"required,dive,required"`
}

func NewQuestionService(repo repository.Questions) *QuestionService {
	return &QuestionService{repo: repo}
}

func (s *QuestionService) GetQuestions() ([]model.Question, error) {
	return s.repo.GetQuestions()
}

func (s *QuestionService) CreateQuestion(input *CreateQuestionInput) (int, error) {
	question := &model.Question{
		Content:   input.Content,
		Type:      input.Type,
		CourseID:  input.CourseID,
		ModuleID:  null.IntFrom(int64(input.ModuleID)),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	questionID, err := s.repo.CreateQuestion(question)
	if err != nil {
		return 0, err
	}

	var answers []model.Answer
	for _, answer := range input.Answers {
		answers = append(answers, model.Answer{
			Content:    answer.Content,
			QuestionID: questionID,
			IsTrue:     answer.IsTrue,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		})
	}
	err = s.repo.CreateAnswers(answers)
	if err != nil {
		return 0, err
	}

	return questionID, nil
}
