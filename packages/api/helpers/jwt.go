package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(id string, username string, role string) (string, error) {
	current_ts := time.Now().Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = id
	claims["iat"] = current_ts
	claims["exp"] = current_ts + 4*60*60
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func validateSignedMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(os.Getenv("JWT_SECRET")), nil
}

type UserJWT struct {
	ID string
}

func VerifyJWT(tokenString string) (bool, *UserJWT) {
	token, err := jwt.Parse(tokenString, validateSignedMethod)

	if err != nil {
		return false, nil
	}
	if !token.Valid {
		return false, nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	id, ok := claims["user_id"].(string)
	if !ok {
		return false, nil
	}

	return true, &UserJWT{
		ID: id,
	}
}
