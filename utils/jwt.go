package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "supersecretkey"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(2 * time.Hour).Unix(),
	})

	return token.SignedString([]byte(SECRET_KEY))

}

func VerifyToken(token string) (int64, error) {

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected SigningMethod.")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return 0, errors.New("Couldnt Parse The Token.")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid Token Claims.")
	}

	if !parsedToken.Valid {
		return 0, errors.New("Invalid Token.")
	}

	userId := int64(claims["userId"].(float64))

	return userId, nil

}
