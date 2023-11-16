package domainuserusecase

import (
	"errors"
	"fmt"
	domainusermodel "order-saga-user/internal/domain/user_auth/model"
	domainuserservice "order-saga-user/internal/domain/user_auth/service"
	drivingportmodel "order-saga-user/internal/port/in/model"
	"order-saga-user/internal/port/out/repository"
)

type UserLoginUseCase struct {
	GetUserAdapter repository.GetUserPort
}

func NewUserLoginUseCase(getUserAdapter repository.GetUserPort) *UserLoginUseCase {
	return &UserLoginUseCase{
		GetUserAdapter: getUserAdapter,
	}
}

func (service *UserLoginUseCase) HandleLoginWithUsername(data drivingportmodel.UserLoginDto) (domainusermodel.User, error) {
	username := data.Username

	user, err := service.GetUserAdapter.GetUser(domainusermodel.User{
		Username: username,
	})

	fmt.Println(user)
	fmt.Println(data)

	if err != nil {
		return domainusermodel.User{}, err
	}

	// verify
	if ok := domainuserservice.PasswordService.Compare(user.Password, data.Password); ok == true {
		return user, nil
	}

	return domainusermodel.User{}, errors.New("Unauthorized")
}
