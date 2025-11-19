package publisher

import (
	"encoding/json"
	"fmt"
	"os"

	models "riskmanagement/models/publisher"

	"github.com/streadway/amqp"
	"gitlab.com/golang-package-library/logger"
)

// PublisherInterface defines the methods that a Publisher should implement
type PublisherInterface interface {
	PublishMessage(dto models.PublishMessageDTO) error
}

// Publisher struct implements the PublisherInterface
type Publisher struct {
	logger     logger.Logger
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// NewPublisher creates a new instance of Publisher
func NewPublisher(
	logger logger.Logger,
) PublisherInterface {
	var err error
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		panic("RABBITMQ_URL is not set")
	}

	conn, err := amqp.Dial(url)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to RabbitMQ: %v", err))
	}
	// Create a channel
	Ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		panic(fmt.Sprintf("failed to open a channel: %v", err))
	}
	return &Publisher{
		Connection: conn,
		Channel:    Ch,
		logger:     logger,
	}
}

// PublishMessage publishes a message to the specified queue
func (r *Publisher) PublishMessage(dto models.PublishMessageDTO) error {
	data, err := json.Marshal(dto)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	err = r.Channel.Publish(
		"",            // exchange
		dto.QueueName, // routing key (queue name)
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	r.logger.Zap.Info("")
	return nil
}
