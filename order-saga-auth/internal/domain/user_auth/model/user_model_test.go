package domainusermodel_test

import (
	"flag"
	"fmt"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	domainusermodel "order-saga-user/internal/domain/user_auth/model"
)

//var _ = Describe("UserModel", func() {
//	var user *domainusermodel.User
//
//	It("When user has enough information", func() {
//		user = &domainusermodel.User{
//			Id:        1,
//			Username:  "lequocthinh12",
//			FirstName: "Thinh",
//			LastName:  "Le",
//			Email:     "quocthinhle2001@gmail.com",
//			Password:  "thinh12345",
//		}
//
//		Expect(user.UserFirstName()).To(Equal("Thinh"))
//		Expect(user.UserLastName()).To(Equal("Le"))
//		Expect(user.UserEmail()).To(Equal("quocthinhle2001@gmail.com"))
//		Expect(user.IsValid()).To(BeTrue())
//
//	})
//
//})
//

var serverAddr, smokeEnv string

func init() {
	flag.StringVar(&serverAddr, "server-addr", "", "Address of the server to smoke-check")
	flag.StringVar(&smokeEnv, "environment", "", "Environment to smoke-check")
}

var _ = ginkgo.DescribeTable("Validate user domain model",
	func(id int, username, firstName, lastName, email, password string) {
		user := &domainusermodel.User{
			Id:        id,
			Username:  username,
			Email:     email,
			LastName:  lastName,
			FirstName: firstName,
			Password:  password,
		}

		g := gomega.NewGomega(ginkgo.Fail)

		g.Expect(serverAddr).NotTo(gomega.BeEmpty())

		fmt.Println("Is environment variables collected?")
		fmt.Println(serverAddr)
		fmt.Println(smokeEnv)
		fmt.Println("============== END ================")

		gomega.Expect(user.UserEmail()).To(gomega.Equal(email))
		gomega.Expect(user.UserLastName()).To(gomega.Equal(lastName))
		gomega.Expect(user.UserFirstName()).To(gomega.Equal(firstName))
		gomega.Expect(user.GetUsername()).To(gomega.Equal(username))
		gomega.Expect(user.IsValid()).To(gomega.BeTrue())
	},
	ginkgo.Entry("Spec 1", 1, "username", "Thinh", "Le", "quocthinhle2001@gmail.com", "thinh12345"),
	ginkgo.Entry("Spec 2", 1, "username", "Thinh", "Le", "quocthinhle2001@gmail.com", "thinh12345"),
	ginkgo.Entry("Spec 3", 1, "username", "Thinh", "Le", "quocthinhle2001@gmail.com", "thinh12345"),
	ginkgo.Entry("Spec 4", 1, "username", "Thinh", "Le", "quocthinhle2001@gmail.com", "thinh12345"),
)
