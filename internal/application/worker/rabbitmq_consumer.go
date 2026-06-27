package worker

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/joaodddev/transaction-processing-engine/internal/domain"
)

func StartConsumer(
	channel *amqp.Channel,
	repository domain.TransactionRepository,
) error {

	messages, err := channel.Consume(
		"transaction-processing",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	go func() {

		for message := range messages {

			var transaction domain.Transaction

			if err := json.Unmarshal(
				message.Body,
				&transaction,
			); err != nil {

				continue
			}

			log.Printf(
				"processing transaction %s",
				transaction.ID,
			)

			transaction.StartProcessing()

			transaction.Approve()

			if err := repository.Save(&transaction); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}
