package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func CalculateAmount(price float64, quantity int64) float64 {
	return price * float64(quantity)
}

var secretKey = []byte("secret-key")

func GenerateToken(u *int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": u,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("failed to create token", err)
		return "", nil
	}
	return tokenString, err
}

func Authorizer(tokenString string) error {
	var secretKey = []byte("secret-key")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}
