package auth

import (
	"time"

	"github.com/codelogydev/core-go/config"
	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte(config.GetEnv("JWT_SECRET", "secret"))

func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ValidateToken(tokenStr string) (int, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !token.Valid {
		return 0, err
	}

	claims := token.Claims.(jwt.MapClaims)
	return int(claims["user_id"].(float64)), nil
}
