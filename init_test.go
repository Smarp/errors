package errors

import (
	"testing"

	"github.com/Sirupsen/logrus"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Company Controller Suit")
}

var _ = BeforeSuite(func() {
	logrus.Info("Test suit starts")
})

var _ = AfterSuite(func() {
	logrus.Info("Test suit ends")
})
