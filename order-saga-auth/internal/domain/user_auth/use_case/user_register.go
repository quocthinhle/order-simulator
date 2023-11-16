package domainuserusecase

import (
	domainusermodel "order-saga-user/internal/domain/user_auth/model"
	drivenport "order-saga-user/internal/port/out/repository"
)

type UserRegisterUseCase struct {
	CreateUserPersistenceAdapter drivenport.CreateUserPort
}

func NewUserRegisterUseCase(createUserPersistenceAdapter drivenport.CreateUserPort) *UserRegisterUseCase {
	return &UserRegisterUseCase{
		CreateUserPersistenceAdapter: createUserPersistenceAdapter,
	}
}

func (service *UserRegisterUseCase) HandleUserRegister(data domainusermodel.User) (interface{}, error) {
	user, err := service.CreateUserPersistenceAdapter.CreateUser(data)

	if err != nil {
		return nil, err
	}

	return user, nil
}
