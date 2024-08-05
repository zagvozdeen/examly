package model

const (
	NewCourseStatus      string = "NEW"
	ActiveCourseStatus   string = "ACTIVE"
	InactiveCourseStatus string = "INACTIVE"
)

var AllCourseStatuses = []string{
	ActiveCourseStatus,
	NewCourseStatus,
	InactiveCourseStatus,
}
