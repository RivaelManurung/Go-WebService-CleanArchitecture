package impl

import (
	"context"
	"fmt"
	"github.com/RivaelManurung/go-webservices-clean-arch/internal/entity"
	"sync"
)

// InMemoryUserRepository adalah implementasi repository pengguna di memori.
type InMemoryUserRepository struct {
	users   map[string]*entity.User // Key: email
	mu      sync.RWMutex
	counter int64
}

// NewInMemoryUserRepository membuat instance baru.
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*entity.User),
	}
}

func (r *InMemoryUserRepository) Save(ctx context.Context, user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.Email]; exists {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}

	r.counter++
	user.ID = r.counter
	r.users[user.Email] = user
	return nil
}

func (r *InMemoryUserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[email]
	if !ok {
		return nil, fmt.Errorf("user with email %s not found", email)
	}
	return user, nil
}
