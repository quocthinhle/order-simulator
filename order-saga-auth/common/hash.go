package appcommon

import (
	"golang.org/x/crypto/bcrypt"
)

const COST = 10

type Hashing struct{}

var hashInstance *Hashing

func init() {
	hashInstance = &Hashing{}
}

func (hash *Hashing) Hash(raw string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(raw), COST)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func (hash *Hashing) Compare(raw, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(raw))

	if err != nil {
		return true
	}

	return false
}

func GetHashInstance() *Hashing {
	return hashInstance
}
