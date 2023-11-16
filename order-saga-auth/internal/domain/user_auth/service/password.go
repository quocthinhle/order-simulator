package domainuserservice

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type PasswordHelper struct {
}

const COST = 10

func (pw *PasswordHelper) Hash(raw string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(raw), COST)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func (pw *PasswordHelper) Compare(hashed, raw string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(raw)); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

var PasswordService *PasswordHelper

func init() {
	PasswordService = &PasswordHelper{}
}
