package middleware

import (
	"cryptoserver/pkg/jwt"
	"errors"
	"net/http"
	"strings"
)

func VerifyAuth(r *http.Request) (string, error) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid authorization header format")
	}

	tokenString := parts[1]

	username,err := jwt.VerifyJwtToken(tokenString)
	if err != nil {
		return "", errors.New("invalid or expired token")
	} 

	return username,nil 

}
