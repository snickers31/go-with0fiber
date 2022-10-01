// ./pkg/utils/jwt_generator.go

package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func GenerateNewAccessToken(uuid uuid.UUID) (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")

	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	claims := jwt.MapClaims{}

	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
	claims["user_id"] = uuid

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return t, nil

}
