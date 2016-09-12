package facebook

type Answer struct {
	Share AnswerShare `json:"share"`
	Error AnswerError `json:"error"`
}
