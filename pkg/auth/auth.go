package auth

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	GenerateToken(userID uint, username string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type authService struct {
	secretKey []byte
}

func NewAuthService(secretKey string) AuthService {
	return &authService{
		secretKey: []byte(secretKey),
	}
}

func (s *authService) GenerateToken(userID uint, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

func (s *authService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})
}