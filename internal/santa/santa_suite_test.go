package santa_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/streadway/amqp"
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

// TODO: cleanup
//var _ = AfterSuite(func() {
//	err = cleanUpTestArtifacts()
//	Expect(err).NotTo(HaveOccurred())
//})
