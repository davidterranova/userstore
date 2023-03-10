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

func TestCreateUser(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	repositories, users := initUserRepositories(t)

	for i := range repositories {
		repo := repositories[i]

		t.Run("it should be possible to create a new user", func(t *testing.T) {
			t.Parallel()

			user := model.NewUser("sam", "smith", "ssmith@userstore.local")

			u, err := repo.CreateUser(ctx, user)
			assert.NoError(t, err)
			assert.Equal(t, *user, *u)
		})

		t.Run("creating an existing user should return user already exists", func(t *testing.T) {
			t.Parallel()

			user := users[0]

			u, err := repo.CreateUser(ctx, user)
			assert.ErrorIs(t, err, adapter.ErrUserAlreadyExist)
			assert.Nil(t, u)
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
