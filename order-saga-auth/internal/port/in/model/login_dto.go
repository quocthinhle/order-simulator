package drivingportmodel

import domainusermodel "order-saga-user/internal/domain/user_auth/model"

type UserLoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserViewPayload struct {
	Username  string
	Email     string
	Id        int
	FirstName string
	LastName  string
}

func (u *UserViewPayload) FromDomainModel(user domainusermodel.User) UserViewPayload {
	u.Id = user.Id
	u.Email = user.Email
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Username = user.Username

	return *u
}
