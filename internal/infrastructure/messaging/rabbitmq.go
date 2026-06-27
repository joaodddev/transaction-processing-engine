package messaging

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/joaodddev/transaction-processing-engine/configs"
)

type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func NewRabbitMQ(
	cfg *configs.Config,
) (*RabbitMQ, error) {

	conn, err := amqp.Dial(
		cfg.RabbitMQURL,
	)

	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()

	if err != nil {
		return nil, err
	}

	err = channel.ExchangeDeclare(
		"transactions",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	_, err = channel.QueueDeclare(
		"transaction-processing",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	err = channel.QueueBind(
		"transaction-processing",
		"transaction.created",
		"transactions",
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	log.Println("RabbitMQ connected")

	return &RabbitMQ{
		Connection: conn,
		Channel:    channel,
	}, nil
}
