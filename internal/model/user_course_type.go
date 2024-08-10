package model

const (
	MarathonUserCourseType string = "MARATHON"
	ModuleUserCourseType   string = "MODULE"
	ErrorUserCourseType    string = "ERROR"
	ExamUserCourseType     string = "EXAM"
)

var AllUserCourseTypes = []string{
	MarathonUserCourseType,
	ModuleUserCourseType,
	ErrorUserCourseType,
	ExamUserCourseType,
}

func GetLabelByCourseType(t string) string {
	switch t {
	case MarathonUserCourseType:
		return "Марафон"
	case ModuleUserCourseType:
		return "Модуль"
	case ErrorUserCourseType:
		return "Ошибки"
	case ExamUserCourseType:
		return "Экзамен"
	default:
		return ""
	}
}

func GetSortByCourseType(t string) int {
	switch t {
	case MarathonUserCourseType:
		return 4
	case ExamUserCourseType:
		return 3
	case ModuleUserCourseType:
		return 2
	case ErrorUserCourseType:
		return 1
	default:
		return 0
	}
}
