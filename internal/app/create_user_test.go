package app_test

import (
	"context"
	"errors"
	"testing"

	"github.com/davidterranova/userstore/internal/adapter"
	"github.com/davidterranova/userstore/internal/app"
	"github.com/davidterranova/userstore/internal/model"
	"github.com/davidterranova/userstore/pkg/xerrors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("with invalid query", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name               string
			query              app.CreateUserCmd
			expectedErrorClass xerrors.Class
		}{
			{
				name:               "with empty command",
				query:              app.CreateUserCmd{},
				expectedErrorClass: xerrors.ClassBadRequest,
			},
			{
				name: "with invalid email",
				query: app.CreateUserCmd{
					FirstName: "john",
					LastName:  "doe",
					Email:     "jdoe",
				},
				expectedErrorClass: xerrors.ClassBadRequest,
			},
		}

		for i := range tests {
			test := tests[i]
			t.Run(test.name, func(t *testing.T) {
				t.Parallel()

				ctr := setup(t)
				handler := app.NewCreateUserHandler(ctr.UserRepository)

				user, err := handler.Handle(ctx, test.query)
				assert.Error(t, err)
				assert.Nil(t, user)

				var classError *xerrors.ClassError
				assert.ErrorAs(t, err, &classError)
				classError, _ = err.(*xerrors.ClassError)
				assert.Equal(t, test.expectedErrorClass, classError.Class())
			})
		}
	})

	t.Run("without repository error", func(t *testing.T) {
		t.Parallel()

		t.Run("it should sanitize inputs", func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			ctr := setup(t)
			handler := app.NewCreateUserHandler(ctr.UserRepository)
			repoUser := model.NewUser("john", "doe", "jdoe@userstore.local")

			ctr.UserRepository.EXPECT().
				CreateUser(ctx, gomock.Any()).
				Times(1).
				Return(repoUser, nil)

			cmd := app.CreateUserCmd{
				FirstName: "John ",
				LastName:  " Doe",
				Email:     "jdoe@userstore.local",
			}
			user, err := handler.Handle(ctx, cmd)
			assert.NoError(t, err)
			assert.Equal(t, user, repoUser)
		})
	})

	t.Run("with already existing user", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		ctr := setup(t)
		handler := app.NewCreateUserHandler(ctr.UserRepository)

		ctr.UserRepository.EXPECT().
			CreateUser(ctx, gomock.Any()).
			Times(1).
			Return(nil, adapter.ErrUserAlreadyExist)

		cmd := app.CreateUserCmd{
			FirstName: "john",
			LastName:  "doe",
			Email:     "jdoe@usertstore.local",
		}
		user, err := handler.Handle(ctx, cmd)
		assert.ErrorIs(t, err, adapter.ErrUserAlreadyExist)
		assert.Nil(t, user)

		var classError *xerrors.ClassError
		assert.ErrorAs(t, err, &classError)
		classError, _ = err.(*xerrors.ClassError)
		assert.Equal(t, xerrors.ClassConflict, classError.Class())
	})

	t.Run("with unknown repository error", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		ctr := setup(t)
		handler := app.NewCreateUserHandler(ctr.UserRepository)

		ctr.UserRepository.EXPECT().
			CreateUser(ctx, gomock.Any()).
			Times(1).
			Return(nil, errors.New("unexpected error"))

		cmd := app.CreateUserCmd{
			FirstName: "john",
			LastName:  "doe",
			Email:     "jdoe@usertstore.local",
		}
		user, err := handler.Handle(ctx, cmd)
		assert.Error(t, err, adapter.ErrUserAlreadyExist)
		assert.Nil(t, user)

		var classError *xerrors.ClassError
		assert.ErrorAs(t, err, &classError)
		classError, _ = err.(*xerrors.ClassError)
		assert.Equal(t, xerrors.ClassInternal, classError.Class())
	})
}
