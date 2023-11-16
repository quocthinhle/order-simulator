package drivingport

import (
	domainusermodel "order-saga-user/internal/domain/user_auth/model"
	drivingportmodel "order-saga-user/internal/port/in/model"
)

type LoginUseCase interface {
	HandleLoginWithUsername(data drivingportmodel.UserLoginDto) (domainusermodel.User, error)
}
