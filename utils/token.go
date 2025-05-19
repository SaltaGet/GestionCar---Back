package utils

import (
	"os"
	"strings"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateUserToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"username":  user.Username,
		"role": user.Role,
	})

	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return t, nil
}

func VerifyToken(tokenString string) (jwt.Claims, error) {
	cleanToken := CleanToken(tokenString)
	token, err := jwt.Parse(cleanToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	return token.Claims, nil
}

func CleanToken(bearerToken string) string {
	const prefix = "Bearer "
	if strings.HasPrefix(bearerToken, prefix) {
		return strings.TrimPrefix(bearerToken, prefix)
	}
	return bearerToken
}

func GenerateWorkplaceToken(workplace *models.Workplace) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": workplace.ID,
	})

	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY_WORKPLACE")))
	if err != nil {
		return "", err
	}

	return t, nil
}

func VerifyWorkplaceToken(tokenString string) (jwt.Claims, error) {
	cleanToken := CleanToken(tokenString)
	token, err := jwt.Parse(cleanToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY_WORKPLACE")), nil
	})
	if err != nil {
		return nil, err
	}

	return token.Claims, nil
}