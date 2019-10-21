package models

type CommResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type TopicResult struct {
	Code   int      `json:"code"`
	Msg    string   `json:"msg"`
	Topics []*Topic `json:"topics"`
}

func NewCommResult() *CommResult {
	return &CommResult{}
}
