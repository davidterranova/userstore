package adapter

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/davidterranova/userstore/internal/model"
	"github.com/google/uuid"
)

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUserAlreadyExist = errors.New("user already exists")
)

type InMemoryUserRepository struct {
	users map[uuid.UUID]model.User
	mutex sync.RWMutex
}

func NewInMemoryUserRepository(users ...*model.User) *InMemoryUserRepository {
	r := &InMemoryUserRepository{
		users: make(map[uuid.UUID]model.User),
	}

	for _, user := range users {
		r.users[user.GetId()] = *user
	}

	return r
}

func (r *InMemoryUserRepository) GetUser(_ context.Context, id uuid.UUID) (*model.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	u, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("%w (%s)", ErrUserNotFound, id.String())
	}

	return &u, nil
}

func (r *InMemoryUserRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	_, err := r.GetUser(ctx, user.GetId())
	if err == nil {
		return nil, ErrUserAlreadyExist
	}

	r.mutex.Lock()
	r.users[user.GetId()] = *user
	r.mutex.Unlock()

	return user, nil
}
