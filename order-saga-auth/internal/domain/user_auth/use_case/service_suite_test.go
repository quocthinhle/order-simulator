package domainuserusecase_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail) // .i.e ginkgo.Fail cause we import using the dot "."

	RunSpecs(t, "Service Suite")
}
