package event

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type EmailPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

func main() {
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	producer, err := NewEventProducer(rabbitConn)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	payload := EmailPayload{
		To:      "kexwaz113@gmail.com",
		Subject: "Verification Code",
		Content: "Verification Code is \n 122345",
	}

	j, err := json.MarshalIndent(&payload, "", "\t")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = producer.Push(string(j), "email.register")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// don't continue until rabbit is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			log.Println("Connected to RabbitMQ!")
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}
