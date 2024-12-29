package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const sercretKey = "secret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(sercretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, error := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // verification de la methode de signature
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(sercretKey), nil
	})

	if error != nil {
		return 0, errors.New("invalid token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token")
	}

	userId := int64(claims["userId"].(float64))
	return userId, nil
}
