package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/RivaelManurung/go-webservices-clean-arch/internal/entity"
	"github.com/RivaelManurung/go-webservices-clean-arch/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// UserUsecase menangani logika bisnis terkait pengguna.
type UserUsecase struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

// NewUserUsecase membuat instance baru dari UserUsecase.
func NewUserUsecase(repo repository.UserRepository, secret string) *UserUsecase {
	return &UserUsecase{
		userRepo:  repo,
		jwtSecret: secret,
	}
}

// Register mendaftarkan pengguna baru.
func (uc *UserUsecase) Register(ctx context.Context, email, password string) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &entity.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	return uc.userRepo.Save(ctx, user)
}

// Login memvalidasi kredensial pengguna dan mengembalikan token JWT.
func (uc *UserUsecase) Login(ctx context.Context, email, password string) (string, error) {
	// Cari pengguna berdasarkan email
	user, err := uc.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	// Bandingkan password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	// Buat token JWT
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(uc.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}