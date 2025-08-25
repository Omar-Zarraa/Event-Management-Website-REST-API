package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//secretKey is the key used to hash the password.
const secretKey = "whoops"

//GenerateAuthToken takes the email and userId as parameters and generates and returns the authentication tken, might also return an error.
func GenerateAuthToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email":  email,
		"UserId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

//VerifyAuthToken takes the token as a parameter and verifies if its valid or not, returns the userId and possibly an error.
func VerifyAuthToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse token")
	}

	if !parsedToken.Valid {
		return 0, errors.New("Invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid token claims")
	}

	// email := claims["Email"].(string)
	userId := int64(claims["UserId"].(float64))
	return userId, nil
}
