package drivingport

import (
	domainusermodel "order-saga-user/internal/domain/user_auth/model"
)

type RegisterUseCase interface {
	HandleUserRegister(data domainusermodel.User) (interface{}, error)
}
