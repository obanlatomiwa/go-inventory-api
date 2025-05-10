package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

func GenerateToken() (string, error) {
	// get jwt secret key
	secret := GetValueFromConfigFile("JWT_SECRET_KEY")

	// get jwt token expire time
	expireMinutesCount, _ := strconv.Atoi(GetValueFromConfigFile("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	claims := jwt.MapClaims{}

	// add expiration time for the token
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expireMinutesCount)).Unix()

	// create a new jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// convert token into a string format
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil

}
