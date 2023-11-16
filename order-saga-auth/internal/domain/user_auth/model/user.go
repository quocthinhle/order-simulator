package domainusermodel

type User struct {
	Id        int
	Username  string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (user *User) IsValid() bool {
	return true
}

func (user *User) UserFirstName() string {
	return user.FirstName
}

func (user *User) UserLastName() string {
	return user.LastName
}

func (user *User) UserEmail() string {
	return user.Email
}

func (user *User) GetUsername() string {
	return user.Username
}
