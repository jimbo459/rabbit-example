package santa_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/streadway/amqp"
	"rabbittest/internal/helpers"
	"testing"
)

var (
	conn *amqp.Connection
	ch   *amqp.Channel
	err  error
)

func TestSanta(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Santa Suite")
}

var _ = AfterSuite(func() {
	err = helpers.CleanUpTestArtifacts()
	Expect(err).NotTo(HaveOccurred())
})
