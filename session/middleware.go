package session

import (
	"errors"
	"fmt"
	"strings"

	"lms_backend/helpers"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func GetTokenFromHeader(c *fiber.Ctx, key, prefix string) (token string, err error) {
	authHeader := c.Get(key)
	token = strings.TrimPrefix(authHeader, prefix+" ")
	if authHeader == "" || token == authHeader {
		err = errors.New("authentication header not present or malformed")
		return
	}
	return
}

func GetClaimsFromToken(token string) (claims *Token, err error) {
	if claims, err = Jwt.ReadToken(token); err == nil {
		return claims, nil
	}
	zap.L().Error("parse token error: ", zap.Error(err))
	return nil, err
}
func TokenMiddleware(c *fiber.Ctx) error {
	tokenString, err := GetTokenFromHeader(c, "Authorization", "Bearer")
	if err != nil {
		return helpers.ResponseUnauthorized(c)
	}

	claims, err := GetClaimsFromToken(tokenString)
	if err != nil {
		return helpers.ResponseUnauthorized(c)
	}

	c.Locals("tokenInfo", claims)
	fmt.Println("auth tokenInfo:", claims)
	return c.Next()
}
