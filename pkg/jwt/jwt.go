package jwt

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(username string) (string,error) {

	claims := jwt.RegisteredClaims {
		IssuedAt: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		Issuer: "cryptoserver",
		Subject: username,
	}

	secretKey := []byte("cryproserver-secret-key")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	return token.SignedString(secretKey)
}

