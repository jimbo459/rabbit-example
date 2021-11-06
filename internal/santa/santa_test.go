package santa_test

import (
	"bytes"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/streadway/amqp"
	"net/http"
	"net/http/httptest"
	"rabbittest/internal/child"
	"rabbittest/internal/santa"
)

var _ = Describe("Santa", func() {
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

	Context("A good child sends a valid request", func() {
		It("works", func() {
			fmt.Println("Running suite")
			var jsonStr = []byte(`{"behaviour":"good", "request":"A toy"}`)
			req, err := http.NewRequest("POST", "/child", bytes.NewBuffer(jsonStr))
			Expect(err).NotTo(HaveOccurred())
			resp := httptest.NewRecorder()
			handler := http.HandlerFunc(child.Producer)
			handler.ServeHTTP(resp, req)

			req, err = http.NewRequest("GET", "/santa", nil)
			Expect(err).NotTo(HaveOccurred())
			resp = httptest.NewRecorder()
			h := http.HandlerFunc(santa.Consumer)
			h.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusOK))
			Expect(resp.Body.String()).To(Equal("Letter received: A toy"))
		})
	})

	Context("A bad child sends a valid request", func() {
		It("works", func() {
			fmt.Println("Running suite")
			var jsonStr = []byte(`{"behaviour":"bad", "request":"A toy"}`)
			req, err := http.NewRequest("POST", "/child", bytes.NewBuffer(jsonStr))
			Expect(err).NotTo(HaveOccurred())
			resp := httptest.NewRecorder()
			handler := http.HandlerFunc(child.Producer)
			handler.ServeHTTP(resp, req)

			req, err = http.NewRequest("GET", "/santa", nil)
			Expect(err).NotTo(HaveOccurred())
			resp = httptest.NewRecorder()
			h := http.HandlerFunc(santa.Consumer)
			h.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusOK))
			Expect(resp.Body.String()).To(Equal(""))
		})
	})

	AfterEach(func() {
		ch.QueuePurge("santasWorkshop", false)
		ch.Close()
		conn.Close()
	})
})
