package model

import (
	"database/sql/driver"
	"errors"
)

type UserRole struct {
	slug string
}

func NewUserRole(s any) (UserRole, error) {
	role, ok := s.(string)
	if !ok {
		return UnknownRole, errors.New("can not assert role to string")
	}
	switch role {
	case GuestRole.slug:
		return GuestRole, nil
	case MemberRole.slug:
		return MemberRole, nil
	case ModeratorRole.slug:
		return ModeratorRole, nil
	case AdminRole.slug:
		return AdminRole, nil
	}

	return UnknownRole, errors.New("unknown role: " + role)
}

var (
	UnknownRole   = UserRole{""}
	GuestRole     = UserRole{"guest"}
	MemberRole    = UserRole{"member"}
	ModeratorRole = UserRole{"moderator"}
	AdminRole     = UserRole{"admin"}
)

func (u UserRole) String() string {
	return u.slug
}

func (u *UserRole) Scan(src any) error {
	r, err := NewUserRole(src)
	if err != nil {
		return err
	}
	*u = r
	return nil
}

func (u UserRole) Value() (driver.Value, error) {
	return u.String(), nil
}
