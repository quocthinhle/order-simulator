package persistence

import (
	appcommon "order-saga-user/common"
	domainusermodel "order-saga-user/internal/domain/user_auth/model"
	drivenportmodel "order-saga-user/internal/port/out/model"
)

type CreateUserAdapter struct {
	AppCtx *appcommon.AppContext
}

func (adapter *CreateUserAdapter) CreateUser(user domainusermodel.User) (domainusermodel.User, error) {
	dto := &drivenportmodel.CreateUserCommandDto{}
	dto.FromDomainModel(user)

	if err := adapter.AppCtx.Db.Create(&dto); err != nil {
		return domainusermodel.User{}, err.Error
	}

	return dto.ToDomainModel(), nil
}

func NewCreateUserAdapter(ctx *appcommon.AppContext) *CreateUserAdapter {
	return &CreateUserAdapter{
		AppCtx: ctx,
	}
}
