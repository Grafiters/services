package rabbitmq

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"riskmanagement/services/tasklistord"

// 	"github.com/streadway/amqp"
// 	"go.uber.org/fx"
// )

// // Interface definition for RabbitMQ operations
// var Module = fx.Provide(
// 	NewRabbitMQ, // Function that returns RabbitMQInterface
// )

// type RabbitMQInterface interface {
// 	ConnectMQ() (err error)
// 	CloseMQ() (err error)
// 	PublishMessage(queueName string, body []byte) error
// 	InsertFromFIle_TasklistORD(queueName string, handler func([]byte) error) error
// }

// // Struct to hold RabbitMQ connection and channel
// type RabbitMQ struct {
// 	Conn    *amqp.Connection
// 	Channel *amqp.Channel
// 	service tasklistord.TasklistORDDefinition
// }

// func NewRabbitMQ(service tasklistord.TasklistORDDefinition) RabbitMQInterface {
// 	url := os.Getenv("RABBITMQ_URL")
// 	if url == "" {
// 		log.Fatalf("RABBITMQ_URL is not set")
// 	}

// 	conn, err := amqp.Dial(url)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
// 	}

// 	ch, err := conn.Channel()
// 	if err != nil {
// 		log.Fatalf("Failed to open a RabbitMQ channel: %s", err)
// 	}

// 	return &RabbitMQ{
// 		Conn:    conn,
// 		Channel: ch,
// 		service: service,
// 	}
// }

// // ConnectMQ establishes a connection and channel with RabbitMQ
// func (rmq *RabbitMQ) ConnectMQ() error {
// 	url := os.Getenv("RABBITMQ_URL")
// 	if url == "" {
// 		return fmt.Errorf("RABBITMQ_URL is not set")
// 	}
// 	conn, err := amqp.Dial(url)
// 	if err != nil {
// 		return err
// 	}

// 	channel, err := conn.Channel()
// 	if err != nil {
// 		conn.Close()
// 		return err
// 	}

// 	rmq.Conn = conn
// 	rmq.Channel = channel
// 	return nil
// }

// func (rmq *RabbitMQ) CloseMQ() error {
// 	if rmq.Channel != nil {
// 		if err := rmq.Channel.Close(); err != nil {
// 			return err
// 		}
// 	}
// 	if rmq.Conn != nil {
// 		if err := rmq.Conn.Close(); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// // PublishMessage publishes a message to the given exchange and routing key
// func (rmq *RabbitMQ) PublishMessage(queueName string, body []byte) error {
// 	if rmq.Channel == nil {
// 		return fmt.Errorf("RabbitMQ channel is not initialized")
// 	}

// 	err := rmq.Channel.Publish(
// 		"",        // exchange
// 		queueName, // routing key
// 		false,     // mandatory
// 		false,     // immediate
// 		amqp.Publishing{
// 			ContentType: "application/json",
// 			Body:        body,
// 		},
// 	)
// 	log.Println("Message published to queue:", queueName)
// 	return err
// }

// // ConsumeMessages sets up a consumer for the given queue
// func (rmq *RabbitMQ) InsertFromFIle_TasklistORD(queueName string, handler func([]byte) error) error {
// 	if rmq.Channel == nil {
// 		return fmt.Errorf("RabbitMQ channel is not initialized")
// 	}

// 	msgs, err := rmq.Channel.Consume(
// 		queueName, // queue
// 		"",        // consumer
// 		true,      // auto-ack
// 		false,     // exclusive
// 		false,     // no-local
// 		false,     // no-wait
// 		nil,       // args
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	go func() {
// 		for msg := range msgs {
// 			if err := handler(msg.Body); err != nil {
// 				// Handle processing error
// 				fmt.Printf("Failed to process message: %v\n", err)
// 			}
// 		}
// 	}()
// 	return nil
// }
