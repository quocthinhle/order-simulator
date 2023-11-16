package appcommon

import (
	_ "crypto/sha256"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var JsonWebTokenInstance *JsonWebToken

func init() {
	publicKey, err := os.ReadFile(os.Getenv("PUBLIC_KEY_ADDR"))

	if err != nil {
		panic(err)
	}

	privateKey, err := os.ReadFile(os.Getenv("PRIVATE_KEY_ADDR"))

	if err != nil {
		panic(err)
	}

	JsonWebTokenInstance = &JsonWebToken{
		encryptPrivateKey: privateKey,
		encryptPublicKey:  publicKey,
	}
}

type JsonWebToken struct {
	encryptPublicKey  []byte
	encryptPrivateKey []byte
	secretKey         []byte
}

type SignOption struct {
	ExpiredIn time.Duration
	Payload   map[string]interface{}
	Alg       jwt.SigningMethod
}

func (j *JsonWebToken) Sign(signOption SignOption) (string, error) {
	expirationTime := time.Now().Add(signOption.ExpiredIn)
	claims := jwt.MapClaims{
		"exp": expirationTime.Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
	}

	for key, val := range signOption.Payload {
		claims[key] = val
	}

	token := jwt.NewWithClaims(signOption.Alg, claims)

	if _, ok := signOption.Alg.(*jwt.SigningMethodHMAC); ok {
		return token.SignedString(j.secretKey)
	}

	if _, ok := signOption.Alg.(*jwt.SigningMethodRSA); ok {
		privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(j.encryptPrivateKey)
		if err != nil {
			return "", err
		}

		return token.SignedString(privateKey)
	}

	return "", errors.New("Cannot sign token")
}

func (j *JsonWebToken) Verify() {}

func (j *JsonWebToken) Decode() {}

func NewJwt() *JsonWebToken {
	return JsonWebTokenInstance
}
