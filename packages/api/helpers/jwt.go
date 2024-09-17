package helpers

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

func validateSignedMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(os.Getenv("JWT_SECRET")), nil
}

type UserJWT struct {
	ID    string
	Email string
}

func VerifyJWT(tokenString string) (bool, *UserJWT) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	_, err := jwt.ParseWithClaims(tokenString, claims, validateSignedMethod)
	if err != nil {
		return false, nil
	}
	email, ok := claims["email"].(string)
	if !ok {
		return false, nil
	}
	id, ok := claims["id"].(string)
	if !ok {
		return false, nil
	}

	return true, &UserJWT{
		Email: email,
		ID:    id,
	}
}
