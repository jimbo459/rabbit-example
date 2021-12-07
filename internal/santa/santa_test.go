package santa_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/streadway/amqp"
	"net/http"
	"net/http/httptest"
	"rabbittest/internal/santa"
)

var _ = Describe("Santa", func() {
	Context("A good child sends a valid request", func() {
		err = setupQueueWithRoutingKey("good")
		Expect(err).NotTo(HaveOccurred())

		req, err := http.NewRequest("GET", "/santa", nil)
		Expect(err).NotTo(HaveOccurred())
		resp := httptest.NewRecorder()
		h := http.HandlerFunc(santa.Consumer)
		h.ServeHTTP(resp, req)

		It("returns status 200", func() {
			Expect(resp.Code).To(Equal(http.StatusOK))
		})
		It("returns the letter request", func() {
			Expect(resp.Body.String()).To(Equal("Letter received: A toy"))
		})
	})

	Context("A bad child sends a valid request", func() {
		err = setupQueueWithRoutingKey("bad")
		Expect(err).NotTo(HaveOccurred())

		req, err := http.NewRequest("GET", "/santa", nil)
		Expect(err).NotTo(HaveOccurred())
		resp := httptest.NewRecorder()
		h := http.HandlerFunc(santa.Consumer)
		h.ServeHTTP(resp, req)

		It("returns status 200", func() {
			Expect(resp.Code).To(Equal(http.StatusOK))
		})
		It("returns no messages", func() {
			Expect(resp.Body.String()).To(Equal(""))
		})

	})
})

func setupQueueWithRoutingKey(key string) error {
	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}

	ch, err = conn.Channel()
	if err != nil {
		return err
	}

	err = ch.ExchangeDeclare(
		"northPole", // name
		"direct",    // type
		true,        // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare(
		"santasWorkshop", // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		return err
	}

	err = ch.QueueBind(
		"santasWorkshop", // queue name
		"good",           // routing key
		"northPole",      // exchange
		false,
		nil)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"northPole", // exchange
		key,         // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("A toy"),
		})
	if err != nil {
		return err
	}

	ch.Close()
	conn.Close()

	return nil
}
