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

	expirationTime := time.Now().Add(jwtDuration)

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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}

	return tokenString, nil
}
func VerifyToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("jwtSecret"), nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("token inválido")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("no se pudieron obtener los claims")
	}

	return claims, nil
}
