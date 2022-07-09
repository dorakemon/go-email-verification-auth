package event

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	conn *amqp.Connection
}

func NewEventProducer(conn *amqp.Connection) (Producer, error) {
	producer := Producer{
		conn: conn,
	}

	err := producer.setup()
	if err != nil {
		return Producer{}, err
	}

	return producer, nil
}

func (e *Producer) setup() error {
	channel, err := e.conn.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()
	return declareExchange(channel)
}

func (e *Producer) Push(payload string, routingKey string) error {
	// channel インスタンスの再生成について
	channel, err := e.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	log.Println("Pushing to channel")

	err = channel.Publish(
		"go-auth-micro", // exchange
		routingKey,      // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
