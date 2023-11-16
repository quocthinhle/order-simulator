package domainuserusecase_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Register", func() {

	BeforeEach(func() {
		fmt.Println("BeforeEach closure")
	})

	Context("User register new account", func() {
		It("Spec 1 ", func() {})
	})

	Context("Context 2", func() {
		It("Spec 2", func() {})
	})

	It("Context 3", func() {})

})
