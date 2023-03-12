//go:build integration

package adapter_test

import (
	"context"
	"sync"
	"testing"

	"github.com/davidterranova/userstore/internal/adapter"
	"github.com/davidterranova/userstore/internal/model"
	"github.com/davidterranova/userstore/pkg/pg"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	seed    = uuid.NewString()
	fixture sync.Once
)

func TestGetUser(t *testing.T) {
	ctx := context.Background()
	repositories, users := initUserRepositories(t)

	for i := range repositories {
		repo := repositories[i]

		t.Run(repo.Name(), func(t *testing.T) {
			t.Parallel()

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
		})
	}
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	repositories, users := initUserRepositories(t)

	for i := range repositories {
		repo := repositories[i]

		t.Run(repo.Name(), func(t *testing.T) {
			t.Run("it should be possible to create a new user", func(t *testing.T) {
				t.Parallel()

				user := model.NewUser("sam", "smith", emailWithSeed(t, "ssmith"))

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
		})
	}
}

func initUserRepositories(t *testing.T) ([]model.UserRepository, []*model.User) {
	t.Helper()

	users := []*model.User{
		model.NewUser("john", "doe", emailWithSeed(t, "jdoe")),
		model.NewUser("dark", "vador", emailWithSeed(t, "dvador")),
		model.NewUser("luke", "skywalker", emailWithSeed(t, "lskywalker")),
	}

	pg := adapter.NewPGUserRepository(pg.TestConnection(t))

	fixture.Do(func() {
		for _, u := range users {
			_, err := pg.CreateUser(context.Background(), u)
			require.NoError(t, err, "user cannot be created %+v", u)
		}
	})

	return []model.UserRepository{pg}, users
}

func emailWithSeed(t *testing.T, email string) string {
	t.Helper()

	return email + "@" + seed + ".local"
}
