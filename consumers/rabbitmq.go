package consumers

import (
	"fmt"
	"os"

	"riskmanagement/consumers/publisher"
	"riskmanagement/consumers/taskassignment"

	"github.com/streadway/amqp"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewConsumers),
	fx.Provide(publisher.NewPublisher),
	fx.Provide(taskassignment.NewReadFileMinioConsumer),
	fx.Provide(taskassignment.NewInserTasklistUkerConsumer),
	fx.Provide(taskassignment.NewInsertLampiranRAPConsumer),
	fx.Provide(taskassignment.NewApprovalTaskAssignmentConsumer),
)

type Consumers struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Consumers  []Consumer
}

type Consumer interface {
	Setup(Channel *amqp.Channel) error
}

func NewConsumers(
	readFileMinioConsumer taskassignment.ReadFileMinioConsumer,
	insertTasklistUkerConsumer taskassignment.InserTasklistUkerConsumer,
	insertTaskAssignmentConsumer taskassignment.InsertLampiranRAPConsumer,
	approvalTaskAssignmentConsumer taskassignment.ApprovalTaskAssignmentConsumer,
) (Consumers, error) {
	// Establish RabbitMQ connection
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		panic("RABBITMQ_URL is not set")
	}

	conn, err := amqp.Dial(url)
	if err != nil {
		return Consumers{}, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	// Create a channel
	Ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return Consumers{}, fmt.Errorf("failed to create channel: %v", err)
	}

	// Return Consumers struct with initialized values
	return Consumers{
		Connection: conn,
		Channel:    Ch,
		Consumers: []Consumer{
			readFileMinioConsumer,
			insertTaskAssignmentConsumer,
			approvalTaskAssignmentConsumer,
			insertTasklistUkerConsumer,
		},
	}, nil
}

func (c Consumers) Connect() error {
	for _, consumer := range c.Consumers {
		err := consumer.Setup(c.Channel)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c Consumers) Close() error {
	if c.Channel != nil {
		if err := c.Channel.Close(); err != nil {
			return err
		}
	}

	if c.Connection != nil {
		if err := c.Connection.Close(); err != nil {
			return err
		}
	}
	return nil
}
