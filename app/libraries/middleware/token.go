package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(userid uint) (string, error) {
	claims := jwt.StandardClaims{
		Subject:   fmt.Sprint(userid),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Contoh: Token berlaku selama 1 hari
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
