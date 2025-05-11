package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"strings"
	"time"
)

type TokenMetadata struct {
	ExpiresAt int64 `json:"expires_at"`
}

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

func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	// verify the token
	token, err := verifyToken(c)

	if err != nil {
		return nil, err
	}

	// extract the token claim data
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		expiresAt := int64(claims["exp"].(float64))

		return &TokenMetadata{
			ExpiresAt: expiresAt,
		}, nil
	}
	return nil, err
}

func CheckToken(c *fiber.Ctx) (bool, error) {
	// get current time
	now := time.Now().Unix()

	claims, err := ExtractTokenMetadata(c)

	if err != nil {
		return false, err
	}

	// get expiration time
	expiresAt := claims.ExpiresAt

	if now > expiresAt {
		return false, err
	}

	return true, nil
}

func extractToken(c *fiber.Ctx) string {
	// get the bearer token from the Authorization header
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")

	if len(token) == 2 {
		return token[0]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtExtractKeyFunc)

	if err != nil {
		return nil, err
	}
	return token, nil
}

func jwtExtractKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(GetValueFromConfigFile("JWT_SECRET_KEY")), nil
}
