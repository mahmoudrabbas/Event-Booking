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

func VerifyToken(token string) error {

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected SigningMethod.")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return errors.New("Couldnt Parse The Token.")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return errors.New("Invalid Token Claims.")
	}

	expFloat := claims["exp"].(float64)

	expInt := int64(expFloat)

	if time.Now().Unix() > expInt {
		return errors.New("Expired Token, Please Signin again.")

	}

	return nil

}
