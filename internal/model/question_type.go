package model

type QuestionType string

const (
	OneAnswerType       string = "ONE_ANSWER"
	MultiplyAnswersType string = "MULTIPLY_ANSWERS"
	InputType           string = "INPUT"
)

var AllQuestionTypes = []string{
	OneAnswerType,
	MultiplyAnswersType,
	InputType,
}
