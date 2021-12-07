package santa

import (
	"fmt"
	"github.com/streadway/amqp"
	"net/http"
	. "rabbittest/internal/helpers"
)

func Consumer(w http.ResponseWriter, req *http.Request) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		"santasWorkshop", // queue
		"",               // consumer
		true,             // auto ack
		false,            // exclusive
		false,            // no local
		false,            // no wait
		nil,              // args
	)
	FailOnError(err, "Failed to register a consumer")
	go func() {
		for d := range msgs {
			fmt.Fprintf(w, "Letter received: %s", d.Body)
		}
	}()

}
