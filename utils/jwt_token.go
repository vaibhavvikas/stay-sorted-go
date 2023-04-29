package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

func GenerateJwtToken(userPid string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_pid"] = userPid
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", errors.Wrap(err, "[GenerateJwtToken]")
	}
	return tokenString, errors.Wrap(err, "[GenerateJwtToken]")
}
