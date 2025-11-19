package models

type QuestionerList struct {
	Id       int64    `json:"id"`
	Question string   `json:"question"`
	Answer   []Answer `json:"answer"`
}

type Question struct {
	Id       int64  `json:"id"`
	Question string `json:"question"`
}

type Answer struct {
	Id         int64   `json:"id"`
	IdQuestion int64   `json:"id_question"`
	Answer     *string `json:"answer"`
	InputType  string  `json:"input_type"`
}
