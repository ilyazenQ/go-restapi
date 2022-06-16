package actions

import (
	"github.com/golang-jwt/jwt"
	"time"
)

func GenerateToken(sign []byte) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()
	claims["user"] = "User"
	claims["authorized"] = true

	tokenString, err := token.SignedString(sign)
	if err != nil {
		return "Error", err
	}
	return tokenString, nil
}
