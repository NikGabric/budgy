package helpers

import (
	"backend/internal/repository"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(hashed), err
}

func CheckPasswordHash(hash string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

func CreateJwt(user repository.GetUserByUsernameRow) (string, error) {

	// TODO: add other claims
	claims := jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"iss":      fmt.Sprintf("http://%s/", os.Getenv("SERVER_HOST")),
		"exp":      time.Now().Add(time.Hour /* TODO: add config */).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(os.Getenv("SERVER_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
