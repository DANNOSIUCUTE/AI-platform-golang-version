package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Khóa bí mật (Đáng lý nên lấy từ biến môi trường os.Getenv)
var JwtSigningKey = []byte("my-super-secret-key-1234")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token sống trong 24 giờ

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSigningKey)

	return tokenString, err
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtSigningKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
