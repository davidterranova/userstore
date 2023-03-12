//go:generate mockgen -destination=user_repository_mock.go -package=model . UserRepository
package model

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	Name() string
	GetUser(ctx context.Context, id uuid.UUID) (*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
}
