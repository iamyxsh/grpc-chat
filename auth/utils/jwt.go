package utils

import (
	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

func GenerateJWT(number string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	//claims["exp"] = string(time.Now().Add(10 * time.Minute).Unix())
	claims["authorized"] = true
	claims["number"] = number

	return token.SignedString(sampleSecretKey)
}
