package utils

import (
	"os"
	"strings"
	"time"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *models.User, tenantID string, memberID string, role *models.Role, permissions *[]string) (string, error) {
	claims := jwt.MapClaims{
        "id":         user.ID,
        "first_name": user.FirstName,
        "last_name":  user.LastName,
        "username":   user.Username,
        "is_admin":   user.IsAdmin,
				"exp":        time.Now().Add(24 * time.Hour).Unix(),
    }

		if tenantID != "" && memberID != "" && role != nil && permissions != nil {
				claims["tenant_id"] = tenantID
				claims["member_id"] = memberID
				claims["role_id"] = role.ID
				claims["role_name"] = role.Name
        claims["permissions"] = permissions
		}

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

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

