package store

import (
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
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
	TagsStore         TagsStore
}

func NewStorage(conn *pgxpool.Pool, log zerolog.Logger) Storage {
	return Storage{
		UsersStore:        &UserStore{conn},
		CoursesStore:      &CourseStore{conn, log},
		ModulesStore:      &ModuleStore{conn, log},
		QuestionsStore:    &QuestionStore{conn, log},
		FilesStore:        &FileStore{conn},
		TestSessionsStore: &TestSessionStore{conn, log},
		UserAnswersStore:  &UserAnswerStore{conn},
		TagsStore:         &TagStore{conn},
	}
}
