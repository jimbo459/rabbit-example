package child_test

import (
	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/streadway/amqp"
	"net/http"
	"net/http/httptest"
	"rabbittest/internal/child"
)

var _ = Describe("Child", func() {
	BeforeEach(func() {
		conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
		Expect(err).NotTo(HaveOccurred())

		ch, err = conn.Channel()
		Expect(err).NotTo(HaveOccurred())

		err = ch.ExchangeDeclare(
			"northPole", // name
			"direct",    // type
			true,        // durable
			false,       // auto-deleted
			false,       // internal
			false,       // no-wait
			nil,         // arguments
		)
		Expect(err).NotTo(HaveOccurred())

		_, err = ch.QueueDeclare(
			"santasWorkshop", // name
			true,             // durable
			false,            // delete when unused
			false,            // exclusive
			false,            // no-wait
			nil,              // arguments
		)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("A child sends a valid request", func() {
		It("should work", func() {
			var jsonStr = []byte(`{"behaviour":"good", "request":"A toy"}`)
			req, err := http.NewRequest("POST", "/child", bytes.NewBuffer(jsonStr))
			Expect(err).NotTo(HaveOccurred())

			resp := httptest.NewRecorder()
			handler := http.HandlerFunc(child.Producer)

			handler.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusCreated))
			Expect(resp.Body.String()).To(Equal("Posted letter to santa"))
		})

	})

	AfterEach(func() {
		ch.QueuePurge("santasWorkshop", false)
		ch.Close()
		conn.Close()
	})
})
