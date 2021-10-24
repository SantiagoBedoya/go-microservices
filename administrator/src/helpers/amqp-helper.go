package helpers

import (
	"os"

	"github.com/streadway/amqp"
)

func SetupAMQP() {
	rabbitmq := os.Getenv("AMQP_URL")

	connection, err := amqp.Dial(rabbitmq)
	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}

	channel, err := connection.Channel()
	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}

	durable, autoDelete := true, false
	internal, noWait := false, false
	err = channel.ExchangeDeclare("events", "topic", durable, autoDelete, internal, noWait, nil)

	if err != nil {
		panic(err)
	}

	durable, exclusive := false, false
	autoDelete, noWait = false, false
	q, err := channel.QueueDeclare("emails", durable, autoDelete, exclusive, noWait, nil)
	if err != nil {
		panic("error declaring the queue: " + err.Error())
	}

	err = channel.QueueBind(q.Name, "#", "events", false, nil)
	if err != nil {
		panic("error binding to the queue: " + err.Error())
	}
}
