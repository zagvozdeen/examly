package store

import (
	"errors"
	"github.com/jackc/pgx/v5"
)

var (
	ErrNotFound = errors.New("resource not found")
	ErrConflict = errors.New("resource already exists")
)

type Storage struct {
	UsersStore        UsersStore
	CoursesStore      CoursesStore
	ModulesStore      ModulesStore
	QuestionsStore    QuestionsStore
	FilesStore        FilesStore
	TestSessionsStore TestSessionsStore
	UserAnswersStore  UserAnswersStore
}

func NewStorage(conn *pgx.Conn) Storage {
	return Storage{
		UsersStore:        &UserStore{conn},
		CoursesStore:      &CourseStore{conn},
		ModulesStore:      &ModuleStore{conn},
		QuestionsStore:    &QuestionStore{conn},
		FilesStore:        &FileStore{conn},
		TestSessionsStore: &TestSessionStore{conn},
		UserAnswersStore:  &UserAnswerStore{conn},
	}
}
