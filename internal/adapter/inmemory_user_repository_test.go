package adapter_test

import (
	"context"
	"testing"

	"github.com/davidterranova/userstore/internal/adapter"
	"github.com/davidterranova/userstore/internal/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	repositories, users := initUserRepositories(t)

	for i := range repositories {
		repo := repositories[i]

		t.Run("it should be possible to fetch an existing user", func(t *testing.T) {
			t.Parallel()

			u := users[0]
			user, err := repo.GetUser(ctx, u.GetId())
			assert.NoError(t, err)
			assert.Equal(t, *u, *user)
		})

		t.Run("if should return user not found if user does not exist", func(t *testing.T) {
			t.Parallel()

			user, err := repo.GetUser(ctx, uuid.New())
			assert.ErrorIs(t, err, adapter.ErrUserNotFound)
			assert.Nil(t, user)
		})
	}
}

func initUserRepositories(t *testing.T) ([]model.UserRepository, []*model.User) {
	t.Helper()

	users := []*model.User{
		model.NewUser("john", "doe", "jdoe@userstore.local"),
		model.NewUser("dark", "vador", "dvador@userstore.local"),
		model.NewUser("luke", "skywalker", "lskywalker@userstore.local"),
	}

	inMemory := adapter.NewInMemoryUserRepository(users...)

	return []model.UserRepository{inMemory}, users
}
