package persistence

import (
	appcommon "order-saga-user/common"
	domainusermodel "order-saga-user/internal/domain/user_auth/model"
	drivenportmodel "order-saga-user/internal/port/out/model"
)

type GetUserPersistence struct {
	AppCtx *appcommon.AppContext
}

func (adapter *GetUserPersistence) GetUser(dto domainusermodel.User) (domainusermodel.User, error) {
	var data drivenportmodel.CreateUserCommandDto

	queryBuilder := adapter.AppCtx.GetDbConnection().Table("users")
	queryBuilder = queryBuilder.Where("username = ?", dto.Username)

	if err := queryBuilder.First(&data).Error; err != nil {
		return domainusermodel.User{}, err
	}

	return (&data).ToDomainModel(), nil
}

func NewGetUserAdapter(ctx *appcommon.AppContext) *GetUserPersistence {
	return &GetUserPersistence{
		AppCtx: ctx,
	}
}
