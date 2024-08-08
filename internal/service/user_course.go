package service

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/Den4ik117/examly/internal/repository"
)

type UserCourseService struct {
	repo repository.UserCourses
}

func NewUserCourseService(repo repository.UserCourses) *UserCourseService {
	return &UserCourseService{repo: repo}
}

func (s *UserCourseService) GetUserCourseByUUID(uuid string) (model.UserCourse, error) {
	course, err := s.repo.GetUserCourseByUUID(uuid)
	if err != nil {
		return course, err
	}

	modules, err := s.repo.GetUserModulesByCourseID(course.ID)
	if err != nil {
		return course, err
	}
	course.Modules = modules

	questions, err := s.repo.GetUserQuestionsByCourseID(course.ID)
	if err != nil {
		return course, err
	}
	course.Questions = questions

	ids := make([]int, len(questions))
	for i, question := range questions {
		ids[i] = question.ID
	}

	answers, err := s.repo.GetUserAnswersByQuestionIDs(ids)
	if err != nil {
		return course, err
	}

	answersByQuestionID := map[int][]model.UserAnswer{}
	for _, answer := range answers {
		answersByQuestionID[answer.QuestionID] = append(answersByQuestionID[answer.QuestionID], answer)
	}

	for i, question := range course.Questions {
		course.Questions[i].Answers = answersByQuestionID[question.ID]
	}

	return course, nil
}
