package service

import (
	"fmt"
	"github.com/Den4ik117/examly/internal/model"
	"github.com/Den4ik117/examly/internal/repository"
	"github.com/Den4ik117/examly/internal/util"
	"github.com/guregu/null/v5"
	"slices"
	"strings"
	"time"
)

type UserQuestionService struct {
	repo repository.UserQuestionsInterface
}

type CheckAnswerInput struct {
	QuestionUUID string `json:"-"`
	AnswerID     int    `json:"answer_id"`
	AnswersIDs   []int  `json:"answers_ids"`
	Input        string `json:"input"`
}

func NewUserQuestionService(repo repository.UserQuestionsInterface) *UserQuestionService {
	return &UserQuestionService{repo: repo}
}

func (s *UserQuestionService) CheckAnswer(input *CheckAnswerInput) (*model.UserQuestion, error) {
	question, err := s.repo.GetUserQuestionByUUID(input.QuestionUUID)
	if err != nil {
		return nil, err
	}
	if question.IsTrue.Valid {
		return nil, fmt.Errorf("question is already answered")
	}

	answers, err := s.repo.GetUserAnswers(question.ID)
	if err != nil {
		return nil, err
	}

	if question.Type == model.OneAnswerType {
		if input.AnswerID == 0 {
			return nil, fmt.Errorf("answer_id is required")
		}

		index := slices.IndexFunc(answers, func(answer model.UserAnswer) bool {
			return answer.ID == input.AnswerID
		})

		if index == -1 {
			return nil, fmt.Errorf("answer_id is invalid")
		}

		answers[index].IsChosen = true
		answers[index].UpdatedAt = time.Now()
		err := s.repo.UpdateUserAnswer(&answers[index])
		if err != nil {
			return nil, err
		}

		question.IsTrue = null.BoolFrom(answers[index].IsTrue)
		question.UpdatedAt = time.Now()
		err = s.repo.UpdateUserQuestion(&question)
		if err != nil {
			return nil, err
		}
	}

	if question.Type == model.MultiplyAnswersType {
		if len(input.AnswersIDs) == 0 {
			return nil, fmt.Errorf("answers_ids is required")
		}

		for i, id := range input.AnswersIDs {
			contains := slices.ContainsFunc(answers, func(answer model.UserAnswer) bool {
				return answer.ID == id
			})

			if !contains {
				return nil, fmt.Errorf("answers_ids[%d] is invalid", i)
			}
		}

		for i, answer := range answers {
			answers[i].IsChosen = slices.ContainsFunc(input.AnswersIDs, func(id int) bool {
				return id == answer.ID
			})
			answers[i].UpdatedAt = time.Now()
			err := s.repo.UpdateUserAnswer(&answers[i])
			if err != nil {
				return nil, err
			}
		}

		question.IsTrue = null.BoolFrom(util.AllFunc(answers, func(answer model.UserAnswer) bool {
			return answer.IsTrue == answer.IsChosen
		}))
		question.UpdatedAt = time.Now()
		err = s.repo.UpdateUserQuestion(&question)
		if err != nil {
			return nil, err
		}
	}

	if question.Type == model.InputType {
		if input.Input == "" {
			return nil, fmt.Errorf("input is required")
		}

		index := slices.IndexFunc(answers, func(answer model.UserAnswer) bool {
			return strings.ToLower(answer.Content) == strings.ToLower(input.Input)
		})

		isTrue := false

		if index == -1 {
			answer := &model.UserAnswer{
				Model: model.Model{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				Content:    input.Input,
				QuestionID: question.ID,
				IsTrue:     false,
				IsChosen:   true,
			}
			err := s.repo.CreateUserAnswer(answer)
			if err != nil {
				return nil, err
			}
		} else {
			answers[index].IsChosen = true
			answers[index].UpdatedAt = time.Now()
			err := s.repo.UpdateUserAnswer(&answers[index])
			if err != nil {
				return nil, err
			}

			isTrue = true
		}

		question.IsTrue = null.BoolFrom(isTrue)
		question.UpdatedAt = time.Now()
		err = s.repo.UpdateUserQuestion(&question)
		if err != nil {
			return nil, err
		}
	}

	course := &model.UserCourse{
		Model: model.Model{
			ID:        question.CourseID,
			UpdatedAt: time.Now(),
		},
		LastQuestionID: null.IntFrom(int64(question.ID)),
	}
	err = s.repo.UpdateUserCourse(course)
	if err != nil {
		return nil, err
	}

	question.Answers = answers

	return &question, nil
}
