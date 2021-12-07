package helpers

import (
	"github.com/streadway/amqp"
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func CleanUpTestArtifacts() error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	ch.QueueDelete("santasWorkshop", false, false, false)
	ch.ExchangeDelete("northPole", false, false)

	ch.Close()
	conn.Close()

	return nil
}
