package application_utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTClaims struct {
	UserID    uuid.UUID `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
}

// CreateToken creates a JWT token with the given claims and key
func CreateToken(claims JWTClaims) (string, error) {
	key := []byte(os.Getenv("JWT_SECRET"))

	mappedClaims := jwt.MapClaims{
		"oid":         claims.UserID.String(),
		"given_name":  claims.FirstName,
		"family_name": claims.LastName,
		"email":       claims.Email,
		"aud":         "https://proxymity.tech/guard",
		"exp":         time.Now().Add(time.Minute * 45).Unix(),
		"iss":         "https://proxymity.tech/guard",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mappedClaims)
	tokenString, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(jwtToken string) (bool, string, error) {
	key := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return false, "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return false, "", err
	}

	return token.Valid, claims["oid"].(string), nil
}
