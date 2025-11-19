package models

type PublishMessageDTO struct {
	QueueName string                 `json:"queueName"`
	Pattern   string                 `json:"pattern"`
	Body      map[string]interface{} `json:"body"`
}
