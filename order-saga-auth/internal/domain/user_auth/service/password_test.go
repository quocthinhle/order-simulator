package domainuserservice_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"math/rand"
	domainuserservice "order-saga-user/internal/domain/user_auth/service"
	"time"
)

func GenerateRandom() string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits + specials
	length := 8
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	str := string(buf)

	return str
}

var _ = Describe("Password", func() {
	Context("Verify password hashing service", func() {
		var rawPassword string
		var hashedPassword string
		var hashHelper *domainuserservice.PasswordHelper
		var hashErr error

		BeforeEach(func() {
			rawPassword = GenerateRandom()
			hashHelper = domainuserservice.PasswordService
			hashedPassword, hashErr = hashHelper.Hash(rawPassword)
		})

		It("Should hash with no error", func() {
			Expect(hashedPassword).NotTo(BeNil())
			Expect(hashErr).To(BeNil())
		})

		It("Should match with raw", func() {
			Expect(hashHelper.Compare(hashedPassword, rawPassword)).To(BeTrue())
		})
	})
})
