// ./pkg/utils/password.go

package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

// $2a$10$zvwH3TKoTeQrE46HdQfUE.OUvOIYk4AkZ3xbBx4AN2vDoKPmRmzi6
func ComparePassword(password, storedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
