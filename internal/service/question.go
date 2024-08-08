package service

import (
	"encoding/json"
	"fmt"
	"github.com/Den4ik117/examly/internal/model"
	"github.com/Den4ik117/examly/internal/repository"
	"github.com/Den4ik117/examly/internal/util"
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
	"os"
	"slices"
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

type ImportQuestionsInput struct {
	CourseID int `json:"course_id" validate:"required"`
	FileID   int `json:"file_id" validate:"required"`
	UserID   int
}

type questionJsonStruct struct {
	Question string   `json:"question"`
	Module   string   `json:"module"`
	Answers  []string `json:"answers"`
	Answer   any      `json:"answer"`
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
	if !util.SomeFunc(input.Answers, func(a Answers) bool {
		return a.IsTrue
	}) {
		return 0, fmt.Errorf("at least one answer must be true")
	}

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

func (s *QuestionService) ImportQuestions(input *ImportQuestionsInput) error {
	f, err := s.repo.GetFileByID(input.FileID)
	if err != nil {
		return err
	}

	bytes, err := os.ReadFile(util.PathToFile(f.Content))
	if err != nil {
		return err
	}

	var questions []questionJsonStruct
	err = json.Unmarshal(bytes, &questions)
	if err != nil {
		return err
	}

	for _, question := range questions {
		uid, err := uuid.NewV7()
		if err != nil {
			return err
		}
		var t string
		var trues []int
		switch v := question.Answer.(type) {
		case float64:
			trues = []int{int(v)}
			t = model.OneAnswerType
		case []any:
			trues = make([]int, len(v))
			for i, i2 := range v {
				d, ok := i2.(float64)
				if !ok {
					return fmt.Errorf("invalid answer type")
				}
				trues[i] = int(d)
			}
			t = model.MultiplyAnswersType
		case nil:
			trues = make([]int, len(question.Answers))
			for i := 0; i < len(question.Answers); i++ {
				trues[i] = i
			}
			t = model.InputType
		default:
			return fmt.Errorf("invalid answer type")
		}

		//m := s.repo.

		q := &model.Question{
			UUID:     uid.String(),
			Content:  question.Question,
			Type:     t,
			Status:   model.NewCourseStatus,
			CourseID: input.CourseID,
			//ModuleID: null.Int{
			//	NullInt64: sql.NullInt64{
			//		Int64: 0,
			//		Valid: false,
			//	},
			//},
			//FileID: null.Int{
			//	NullInt64: sql.NullInt64{
			//		Int64: 0,
			//		Valid: false,
			//	},
			//},
			UserID:    input.UserID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		id, err := s.repo.CreateQuestion(q)
		if err != nil {
			return err
		}
		q.ID = id

		answers := make([]model.Answer, len(question.Answers))
		for i, answer := range question.Answers {
			answers[i] = model.Answer{
				Content:    answer,
				QuestionID: q.ID,
				IsTrue:     slices.Contains(trues, i),
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			}
		}
		err = s.repo.CreateAnswers(answers)
		if err != nil {
			return err
		}
	}

	//log.Println(questions)

	return nil
}
