package models

type HttpResponse[T any] struct {
	StatusCode int      `json:"status_code"`
	Message    string   `json:"message"`
	Data       T        `json:"data"`
	Errors     []string `json:"errors"`
}

type HttpResResponse struct {
	StatusCode int      `json:"status_code"`
	Message    string   `json:"message"`
	Data       any      `json:"data"`
	Errors     []string `json:"errors"`
}
