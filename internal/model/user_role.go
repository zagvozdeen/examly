package model

type userRole string

const (
	AdminRole      userRole = "ADMIN"
	ManagerRole    userRole = "MANAGER"
	SimpleUserRole userRole = "USER"
)

type UserRole interface {
	Value() userRole
}

func (r userRole) Value() userRole {
	return r
}
