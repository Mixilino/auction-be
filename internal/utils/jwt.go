package utils

import (
	"auction-be/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

type myCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int, username string) (string, error) {
	claims := myCustomClaims{
		username,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    username,
			ID:        strconv.Itoa(userId),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(config.GetJWTSecret()))
	if err != nil {
		return "", err
	}
	return ss, nil
}
