package child_test

import (
	"github.com/streadway/amqp"
	"rabbittest/internal/helpers"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	conn *amqp.Connection
	ch   *amqp.Channel
	err  error
)

func TestChild(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Child Suite")
}

var _ = AfterSuite(func() {
	err = helpers.CleanUpTestArtifacts()
	Expect(err).NotTo(HaveOccurred())
})
