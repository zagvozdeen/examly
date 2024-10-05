package store

import "testing"

func NewMockStorage(t *testing.T) Storage {
	return Storage{
		UsersStore:        NewMockUsersStore(t),
		CoursesStore:      NewMockCoursesStore(t),
		ModulesStore:      NewMockModulesStore(t),
		QuestionsStore:    NewMockQuestionsStore(t),
		FilesStore:        NewMockFilesStore(t),
		TestSessionsStore: NewMockTestSessionsStore(t),
		UserAnswersStore:  NewMockUserAnswersStore(t),
	}
}
