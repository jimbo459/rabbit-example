package child

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"net/http"
	. "rabbittest/internal/helpers"
)

type letter struct {
	Behaviour string
	Request   string
}

func Producer(w http.ResponseWriter, req *http.Request) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	var childLetter letter
	err = json.NewDecoder(req.Body).Decode(&childLetter)
	if err != nil {
		http.Error(w, "Request body invalid", http.StatusBadRequest)
		return
	}

	err = ch.Publish(
		"northPole",           // exchange
		childLetter.Behaviour, // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(childLetter.Request),
		})
	FailOnError(err, "Failed to publish a message")

	w.WriteHeader(http.StatusCreated)

	fmt.Fprintf(w, "Posted letter to santa")
}
