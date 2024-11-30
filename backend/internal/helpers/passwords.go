package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(hashed), err
}
