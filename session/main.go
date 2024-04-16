package session

import (
	"os"

	"github.com/golang-jwt/jwt"
)

func InitGjwt(private, public string) error {
	var (
		key    []byte
		result JWT
	)
	key, err := os.ReadFile(private)
	if err != nil {
		return err
	}
	result.private, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(key))
	if err != nil {
		return err
	}
	key, err = os.ReadFile(public)
	if err != nil {
		return err
	}
	result.pub, err = jwt.ParseRSAPublicKeyFromPEM([]byte(key))

	Jwt = &result
	return err
}
