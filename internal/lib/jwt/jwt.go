package jwtlib

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aidosgal/gust/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

func NewToken(user model.User, app model.App, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.Id
	claims["phone"] = user.Phone
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string, secret string) (userID int, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("Parsed Claims: %+v\n", claims) // Debugging
		uid, ok := claims["uid"].(float64)
		if !ok {
			return 0, errors.New("user_id claim not found or invalid")
		}

		return int(uid), nil
	}

	return 0, errors.New("invalid token or claims")
}

func GetToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("Missing Authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", fmt.Errorf("Invalid Authorization header format")
	}

	token := parts[1]
	return token, nil
}
