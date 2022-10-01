// ./pkg/utils/jwt_parser.go

package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type TokenMetadata struct {
	Expires int64
}

func ExtractTokenMetaData(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	fmt.Println(token)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		expires := int64(claims["exp"].(float64))

		return &TokenMetadata{
			Expires: expires,
		}, nil
	}

	return nil, err

}
func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}

func extractToken(c *fiber.Ctx) (string, error) {
	bearToken := c.Get("Authorization")

	onlyToken := strings.Split(bearToken, " ")

	if len(onlyToken) == 2 {
		return onlyToken[1], nil
	}
	return "", fmt.Errorf("Unauthorized")
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString, err := extractToken(c)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenString, jwtKeyFunc)

	if err != nil {
		return nil, err
	}

	return token, nil
}
