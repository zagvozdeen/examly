package enum

import (
	"database/sql/driver"
	"errors"
)

type QuestionType struct {
	slug string
}

func NewQuestionType(s any) (QuestionType, error) {
	role, ok := s.(string)
	if !ok {
		return UnknownQuestionType, errors.New("can not assert role to string")
	}
	switch role {
	case SingleChoiceQuestionType.slug:
		return SingleChoiceQuestionType, nil
	case MultipleChoiceQuestionType.slug:
		return MultipleChoiceQuestionType, nil
	case PlaintextQuestionType.slug:
		return PlaintextQuestionType, nil
	}

	return UnknownQuestionType, errors.New("unknown role: " + role)
}

var (
	UnknownQuestionType        = QuestionType{""}
	SingleChoiceQuestionType   = QuestionType{"single_choice"}
	MultipleChoiceQuestionType = QuestionType{"multiple_choice"}
	PlaintextQuestionType      = QuestionType{"plaintext"}
)

func (u QuestionType) String() string {
	return u.slug
}

func (u *QuestionType) Scan(src any) error {
	r, err := NewQuestionType(src)
	if err != nil {
		return err
	}
	*u = r
	return nil
}

func (u QuestionType) Value() (driver.Value, error) {
	return u.String(), nil
}

func (u QuestionType) MarshalJSON() ([]byte, error) {
	return []byte("\"" + u.String() + "\""), nil
}
