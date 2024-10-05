package enum

import (
	"database/sql/driver"
	"errors"
)

type Status struct {
	slug string
}

func NewStatus(s any) (Status, error) {
	role, ok := s.(string)
	if !ok {
		return UnknownStatus, errors.New("can not assert role to string")
	}
	switch role {
	case CreatedStatus.slug:
		return CreatedStatus, nil
	case ActiveStatus.slug:
		return ActiveStatus, nil
	case InactiveStatus.slug:
		return InactiveStatus, nil
	}

	return UnknownStatus, errors.New("unknown role: " + role)
}

var (
	UnknownStatus  = Status{""}
	CreatedStatus  = Status{"created"}
	ActiveStatus   = Status{"active"}
	InactiveStatus = Status{"inactive"}
)

func (u Status) String() string {
	return u.slug
}

func (u *Status) Scan(src any) error {
	r, err := NewStatus(src)
	if err != nil {
		return err
	}
	*u = r
	return nil
}

func (u Status) Value() (driver.Value, error) {
	return u.String(), nil
}
