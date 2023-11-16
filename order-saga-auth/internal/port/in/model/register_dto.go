package drivingportmodel

import domainusermodel "order-saga-user/internal/domain/user_auth/model"

type UserRegisterDto struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

func (dto UserRegisterDto) ToUserDomainModel() domainusermodel.User {
	return domainusermodel.User{
		Username:  dto.Username,
		LastName:  dto.LastName,
		FirstName: dto.FirstName,
		Password:  dto.Password,
		Email:     dto.Email,
	}
}
