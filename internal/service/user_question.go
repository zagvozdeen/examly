package service

import (
	"database/sql"
	"errors"
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
	repo        repository.UserQuestionsInterface
	courses     repository.Courses
	userCourses repository.UserCourses
}

type CheckAnswerInput struct {
	QuestionUUID string `json:"-"`
	UserID       int    `json:"-"`
	AnswerID     int    `json:"answer_id"`
	AnswersIDs   []int  `json:"answers_ids"`
	Input        string `json:"input"`
}

func NewUserQuestionService(
	repo repository.UserQuestionsInterface,
	courses repository.Courses,
	userCourses repository.UserCourses,
) *UserQuestionService {
	return &UserQuestionService{
		repo:        repo,
		courses:     courses,
		userCourses: userCourses,
	}
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

	isTrue := false

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

		isTrue = answers[index].IsTrue
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

		isTrue = util.AllFunc(answers, func(answer model.UserAnswer) bool {
			return answer.IsTrue == answer.IsChosen
		})
	}

	if question.Type == model.InputType {
		if input.Input == "" {
			return nil, fmt.Errorf("input is required")
		}

		index := slices.IndexFunc(answers, func(answer model.UserAnswer) bool {
			return strings.ToLower(answer.Content) == strings.ToLower(input.Input)
		})

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
			if err := s.repo.CreateUserAnswer(answer); err != nil {
				return nil, err
			}
		} else {
			answers[index].IsChosen = true
			answers[index].UpdatedAt = time.Now()
			if err := s.repo.UpdateUserAnswer(&answers[index]); err != nil {
				return nil, err
			}

			isTrue = true
		}
	}

	question.IsTrue = null.BoolFrom(isTrue)
	question.UpdatedAt = time.Now()
	if err = s.repo.UpdateUserQuestion(&question); err != nil {
		return nil, err
	}

	if !isTrue {
		if err = s.addQuestionToCourseWithErrors(input, question, answers); err != nil {
			return nil, err
		}
	}

	course := &model.UserCourse{}
	course.ID = question.CourseID
	course.UpdatedAt = time.Now()
	course.LastQuestionID = null.IntFrom(int64(question.ID))
	if err = s.repo.UpdateUserCourse(course); err != nil {
		return nil, err
	}

	question.Answers = answers

	return &question, nil
}

func (s *UserQuestionService) addQuestionToCourseWithErrors(
	input *CheckAnswerInput, question model.UserQuestion, answers []model.UserAnswer,
) error {
	course, err := s.userCourses.GetUserCourseByTypeAndUserID(model.ErrorUserCourseType, input.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			original, err := s.userCourses.GetUserCourseByID(question.CourseID)
			if err != nil {
				return err
			}
			course.UUID = util.GenerateUUID()
			course.Name = original.Name
			course.Type = model.ErrorUserCourseType
			course.UserID = input.UserID
			course.CourseID = original.CourseID
			course.CreatedAt = time.Now()
			course.UpdatedAt = time.Now()
			if err = s.courses.CreateUserCourse(&course); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	newQ := &model.UserQuestion{
		Model: model.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UUID:        util.GenerateUUID(),
		Content:     question.Content,
		Explanation: question.Explanation,
		Type:        question.Type,
		Sort:        0,
		CourseID:    course.ID,
		QuestionID:  question.QuestionID,
		ModuleID:    question.ModuleID,
		FileID:      question.FileID,
	}
	if err = s.userCourses.CreateUserQuestion(newQ); err != nil {
		return err
	}

	newA := make([]model.UserAnswer, len(answers))
	for i, answer := range answers {
		newA[i] = model.UserAnswer{
			Model: model.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Content:    answer.Content,
			QuestionID: question.ID,
			IsTrue:     answer.IsTrue,
			IsChosen:   false,
			Sort:       answer.Sort,
		}
	}
	if err = s.courses.CreateUserAnswers(newA); err != nil {
		return err
	}

	return nil
}
