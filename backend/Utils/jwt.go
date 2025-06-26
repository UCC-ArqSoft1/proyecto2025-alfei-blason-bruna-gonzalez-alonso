package Utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	jwtDuration = time.Hour * 24
	jwtSecret   = "jwtSecret"
)

type CustomClaims struct {
	UserID  int  `json:"user_id"`
	IsAdmin bool `json:"is_admin"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int, isAdmin bool) (string, error) {
	// Setear expiracion
	expirationTime := time.Now().Add(jwtDuration)

	// Construir los claims
	claims := CustomClaims{
		UserID:  userID,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "backend",
			Subject:   "auth",
			ID:        fmt.Sprintf("%d", userID),
		},
	}

	// Crear el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}

	return tokenString, nil
}
