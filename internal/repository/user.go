package repository

import (
	"context"
	"github.com/RivaelManurung/go-webservices-clean-arch/internal/entity"
)

// UserRepository adalah interface untuk data pengguna.
type UserRepository interface {
	Save(ctx context.Context, user *entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}