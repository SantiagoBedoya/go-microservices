package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/streadway/amqp"
	gomail "gopkg.in/mail.v2"
)

func main() {
	fmt.Println("running emails microservice -- listen emails queue")
	rabbitmq := os.Getenv("AMQP_URL")

	connection, err := amqp.Dial(rabbitmq)
	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}

	channel, err := connection.Channel()
	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}

	msgs, err := channel.Consume("emails", "", false, false, false, false, nil)

	if err != nil {
		panic("error consuming the queue: " + err.Error())
	}

	host := os.Getenv("EMAIL_HOST")
	port, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	username := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")

	d := gomail.NewDialer(host, port, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	for msg := range msgs {

		var data map[string]string
		json.Unmarshal(msg.Body, &data)
		msg.Ack(false)

		go func() {
			m := gomail.NewMessage()
			m.SetHeader("From", username)
			m.SetHeader("To", data["email"])
			m.SetHeader("Subject", data["subject"])
			m.SetBody("text/plain", data["message"])

			if err := d.DialAndSend(m); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("sended to " + data["email"])
			}

		}()
	}

	defer connection.Close()
}
