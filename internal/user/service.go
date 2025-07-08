package user

import (
	"go-webservices-clean-arch/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo    UserRepository
	authSvc auth.AuthService
}

func NewUserService(repo UserRepository, authSvc auth.AuthService) UserService {
	return &userService{
		repo:    repo,
		authSvc: authSvc,
	}
}

func (s *userService) Register(user *User) (*User, string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}
	user.Password = string(hashedPassword)

	if err := s.repo.Create(user); err != nil {
		return nil, "", err
	}

	token, err := s.authSvc.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *userService) Login(username, password string) (string, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	token, err := s.authSvc.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
func (s *userService) GetProfile(id  uint) (*User, error) {
		return s.repo.FindByID(id)
}