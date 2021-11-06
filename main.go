package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"net/http"
	"rabbittest/internal/child"
	. "rabbittest/internal/helpers"
	"rabbittest/internal/santa"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"northPole", // name
		"direct",    // type
		true,        // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	FailOnError(err, "Failed to declare an exchange")

	_, err = ch.QueueDeclare(
		"santasWorkshop", // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	http.HandleFunc("/child", child.Producer)
	http.HandleFunc("/santa", santa.Consumer)

	fmt.Printf("Listening on address localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
