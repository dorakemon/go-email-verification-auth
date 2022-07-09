package event

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn *amqp.Connection
}

func NewEventConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}

	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

func (e *Consumer) setup() error {
	channel, err := e.conn.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()
	return declareExchange(channel)
}

type handleFunc func([]byte)

func (e *Consumer) ListenRegisterEmail(handleFunc handleFunc) error {
	channel, err := e.conn.Channel()
	if err != nil {
		return err
	}

	q, err := channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	err = channel.QueueBind(
		q.Name,           // queue name
		"email.register", // routing key
		"go-auth-micro",  // exchange
		false,
		nil)
	if err != nil {
		return err
	}

	msgs, err := channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			go handleFunc(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

	return nil
}

// func (e *Consumer) Push(payload string, routingKey string) error {
// 	// channel インスタンスの再生成について
// 	channel, err := e.conn.Channel()
// 	if err != nil {
// 		return err
// 	}
// 	defer channel.Close()

// 	log.Println("Pushing to channel")

// 	err = channel.Publish(
// 		"go-auth-micro", // exchange
// 		routingKey,      // routing key
// 		false,           // mandatory
// 		false,           // immediate
// 		amqp.Publishing{
// 			ContentType: "text/plain",
// 			Body:        []byte(payload),
// 		},
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
