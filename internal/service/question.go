package service

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/Den4ik117/examly/internal/repository"
	"github.com/google/uuid"
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
	Content     string    `json:"content" validate:"required"`
	Explanation string    `json:"explanation" validate:""`
	Type        string    `json:"type" validate:"required"`
	CourseID    int       `json:"course_id" validate:"required"`
	FileID      int       `json:"file_id" validate:""`
	ModuleID    int       `json:"module_id" validate:""`
	Answers     []Answers `json:"answers" validate:"required,dive,required"`
}

func NewQuestionService(repo repository.Questions) *QuestionService {
	return &QuestionService{repo: repo}
}

func (s *QuestionService) GetQuestions() ([]model.Question, error) {
	return s.repo.GetQuestions()
}

func (s *QuestionService) GetQuestionsByUserID(id int) ([]model.Question, error) {
	return s.repo.GetQuestionsByUserID(id)
}

func (s *QuestionService) CreateQuestion(user *model.User, input *CreateQuestionInput) (int, error) {
	uid, err := uuid.NewV7()
	if err != nil {
		return 0, err
	}

	question := &model.Question{
		UUID:        uid.String(),
		Content:     input.Content,
		Explanation: null.NewString(input.Explanation, input.Explanation != ""),
		Type:        input.Type,
		Status:      model.NewCourseStatus,
		CourseID:    input.CourseID,
		ModuleID:    null.NewInt(int64(input.ModuleID), input.ModuleID != 0),
		FileID:      null.NewInt(int64(input.FileID), input.FileID != 0),
		UserID:      user.ID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
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
