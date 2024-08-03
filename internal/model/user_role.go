package model

import (
	"database/sql/driver"
)

type UserRole string

const (
	AdminRole      UserRole = "ADMIN"
	ManagerRole    UserRole = "MANAGER"
	SimpleUserRole UserRole = "USER"
	GuestRole      UserRole = "GUEST"
)

func (u *UserRole) Scan(src interface{}) error {
	*u = UserRole(src.(string))
	return nil
}

func (u UserRole) Value() (driver.Value, error) {
	return string(u), nil
}
