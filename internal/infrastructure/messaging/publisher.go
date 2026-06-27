package messaging

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	channel *amqp.Channel
}

func NewPublisher(
	channel *amqp.Channel,
) *Publisher {

	return &Publisher{
		channel: channel,
	}
}

func (p *Publisher) Publish(
	body []byte,
) error {

	return p.channel.PublishWithContext(
		context.Background(),
		"transactions",
		"transaction.created",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
