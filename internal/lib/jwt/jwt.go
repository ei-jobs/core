package jwtlib

import (
	"errors"
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

func ParseToken(tokenString string, secret string) (userID int, appID int, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uid, ok := claims["uid"].(float64)
		if !ok {
			return 0, 0, errors.New("user_id claim not found or invalid")
		}

		aid, ok := claims["app_id"].(float64)
		if !ok {
			return 0, 0, errors.New("app_id claim not found or invalid")
		}

		return int(uid), int(aid), nil
	}

	return 0, 0, errors.New("invalid token or claims")
}
