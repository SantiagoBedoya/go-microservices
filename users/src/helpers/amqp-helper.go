package helpers

import (
	"encoding/json"
	"os"
	"time"

	"github.com/streadway/amqp"
)

type Data struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func PublishToEmailQueque(d Data) {
	rabbitmq := os.Getenv("AMQP_URL")

	connection, err := amqp.Dial(rabbitmq)
	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}

	channel, err := connection.Channel()
	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}

	data, err := json.Marshal(d)
	if err != nil {
		panic("could not marshal json")
	}

	msg := amqp.Publishing{
		DeliveryMode: 1,
		Timestamp:    time.Now(),
		ContentType:  "application/json",
		Body:         data,
	}
	channel.Publish("events", "emails", false, false, msg)
	defer connection.Close()
}
