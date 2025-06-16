package common

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func ParseJWT(tokenStr string) (jwt.MapClaims, error) {

	var secretKey = []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}
	return claims, nil
}
