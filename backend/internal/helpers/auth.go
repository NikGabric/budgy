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
		"sub":      int32(user.ID),
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

func DecodeJwt(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(os.Getenv("SERVER_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return nil, fmt.Errorf("Token has expired")
			}
		}

		return claims, nil
	}

	return nil, fmt.Errorf("Invalid token")
}

func GetClaimFromJwt(tokenStr, claimKey string) (interface{}, error) {
	claims, err := DecodeJwt(tokenStr)
	if err != nil {
		return nil, err
	}

	claim, exists := claims[claimKey]
	if !exists {
		return nil, fmt.Errorf("claim %s not found", claimKey)
	}

	return claim, nil
}
