package child_test

import (
	"github.com/streadway/amqp"
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
