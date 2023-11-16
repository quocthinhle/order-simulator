package drivenportmodel

import (
	domainusermodel "order-saga-user/internal/domain/user_auth/model"
	domainuserservice "order-saga-user/internal/domain/user_auth/service"
)

type CreateUserCommandDto struct {
	Id        int    `gorm:"column:id;"`
	Username  string `gorm:"column:username"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
}

func (CreateUserCommandDto) TableName() string {
	return "users"
}

func (dto *CreateUserCommandDto) ToDomainModel() domainusermodel.User {
	return domainusermodel.User{
		Username:  dto.Username,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  dto.Password,
		Id:        dto.Id,
	}
}

func (dto *CreateUserCommandDto) FromDomainModel(user domainusermodel.User) {
	hashedPassword, err := domainuserservice.PasswordService.Hash(user.Password)

	if err != nil {
		panic("Handle error later :))")
	}

	dto.Id = user.Id
	dto.LastName = user.LastName
	dto.FirstName = user.FirstName
	dto.Email = user.Email
	dto.Password = hashedPassword
	dto.Username = user.Username
}
