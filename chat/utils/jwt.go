package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

func VerifyJWT(token string) (string, bool) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's an error with the signing method")
		}
		return sampleSecretKey, nil

	})

	var number string

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {

		number = claims["number"].(string)
	} else {
		fmt.Println(err)
	}

	if err != nil {
		return "", false
	}

	return number, true
}
