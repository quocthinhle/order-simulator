package repository

import (
	domainusermodel "order-saga-user/internal/domain/user_auth/model"
)

type CreateUserPort interface {
	CreateUser(user domainusermodel.User) (domainusermodel.User, error)
}
