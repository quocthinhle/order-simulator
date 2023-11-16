package repository

import (
	domainusermodel "order-saga-user/internal/domain/user_auth/model"
)

type GetUserPort interface {
	GetUser(dto domainusermodel.User) (domainusermodel.User, error)
}
