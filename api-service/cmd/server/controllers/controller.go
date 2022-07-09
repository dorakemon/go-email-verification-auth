package controllers

import amqp "github.com/rabbitmq/amqp091-go"

type Config struct {
	Conn *amqp.Connection
}
