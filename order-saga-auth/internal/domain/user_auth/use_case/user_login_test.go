package domainuserusecase_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	domainusermodel "order-saga-user/internal/domain/user_auth/model"
	domainuserservice "order-saga-user/internal/domain/user_auth/service"
	domainuserusecase "order-saga-user/internal/domain/user_auth/use_case"
	drivingportmodel "order-saga-user/internal/port/in/model"
	"order-saga-user/internal/port/out/repository"
)

type GetUserPortMock struct{}

func (g *GetUserPortMock) GetUser(dto domainusermodel.User) (domainusermodel.User, error) {
	hash, err := domainuserservice.PasswordService.Hash("thinh12345")

	if err != nil {
		panic(err)
	}

	fmt.Println(hash)

	return domainusermodel.User{
		Username:  dto.Username,
		Email:     dto.Email,
		Password:  hash,
		Id:        dto.Id,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	}, nil
}

var _ = Describe("LoginUsecase", func() {
	Context("Login Usecase", func() {
		// var myCall *gomock.Controller
		var loginUsecase *domainuserusecase.UserLoginUseCase
		var getUserPersistenceAdapter repository.GetUserPort

		BeforeEach(func() {
			// Somehow we mock getUserAdapter
			// myCall = gomock.NewController(GinkgoT())
			getUserPersistenceAdapter = &GetUserPortMock{}
			loginUsecase = domainuserusecase.NewUserLoginUseCase(getUserPersistenceAdapter)
		})

		It("Should login correctly when provide valid credentials", func() {
			user, err := loginUsecase.HandleLoginWithUsername(drivingportmodel.UserLoginDto{
				Username: "lequocthinh12",
				Password: "thinh12345",
			})

			Expect(user).NotTo(BeNil())
			Expect(err).To(BeNil())
		})
	})
})
