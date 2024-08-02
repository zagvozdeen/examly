package model

type courseType string

const (
	OneAnswerType       courseType = "ONE_ANSWER"
	MultiplyAnswersType courseType = "MULTIPLY_ANSWERS"
	InputType           courseType = "INPUT"
)

type CourseType interface {
	Value() courseType
}

func (r courseType) Value() courseType {
	return r
}
