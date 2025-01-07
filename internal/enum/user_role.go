package enum

import (
	"database/sql/driver"
	"errors"
)

type UserRole struct {
	slug  string
	level int8
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
	case ReferralRole.slug:
		return ReferralRole, nil
	case CompanyRole.slug:
		return CompanyRole, nil
	case ModeratorRole.slug:
		return ModeratorRole, nil
	case AdminRole.slug:
		return AdminRole, nil
	}

	return UnknownRole, errors.New("unknown role: " + role)
}

var (
	UnknownRole   = UserRole{"", 0}
	GuestRole     = UserRole{"guest", 1}
	MemberRole    = UserRole{"member", 2}
	ReferralRole  = UserRole{"referral", 3}
	CompanyRole   = UserRole{"company", 4}
	ModeratorRole = UserRole{"moderator", 5}
	AdminRole     = UserRole{"admin", 6}
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

func (u UserRole) Level() int8 {
	return u.level
}

func (u UserRole) MarshalJSON() ([]byte, error) {
	return []byte("\"" + u.String() + "\""), nil
}
