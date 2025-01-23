package enum

import (
	"database/sql/driver"
	"errors"
)

type TestSessionType struct {
	slug string
}

func NewTestSessionType(s any) (TestSessionType, error) {
	role, ok := s.(string)
	if !ok {
		return UnknownTestSessionType, errors.New("can not assert test session type to string")
	}
	switch role {
	case SelectionSystemTestSessionType.slug:
		return SelectionSystemTestSessionType, nil
	case MarathonTestSessionType.slug:
		return MarathonTestSessionType, nil
	case MistakeTestSessionType.slug:
		return MistakeTestSessionType, nil
	case ModuleTestSessionType.slug:
		return ModuleTestSessionType, nil
	case ExamTestSessionType.slug:
		return ExamTestSessionType, nil
	}

	return UnknownTestSessionType, errors.New("unknown test session type: " + role)
}

var (
	UnknownTestSessionType         = TestSessionType{""}
	SelectionSystemTestSessionType = TestSessionType{"selection-system"}
	MarathonTestSessionType        = TestSessionType{"marathon"}
	MistakeTestSessionType         = TestSessionType{"mistake"}
	ModuleTestSessionType          = TestSessionType{"module"}
	ExamTestSessionType            = TestSessionType{"exam"}
)

func (u TestSessionType) String() string {
	return u.slug
}

func (u *TestSessionType) Scan(src any) error {
	r, err := NewTestSessionType(src)
	if err != nil {
		return err
	}
	*u = r
	return nil
}

func (u TestSessionType) Value() (driver.Value, error) {
	return u.String(), nil
}

func (u TestSessionType) MarshalJSON() ([]byte, error) {
	return []byte("\"" + u.String() + "\""), nil
}
