package app_test

import (
	"context"
	"errors"
	"testing"

	"github.com/davidterranova/userstore/internal/adapter"
	"github.com/davidterranova/userstore/internal/app"
	"github.com/davidterranova/userstore/internal/domain"
	"github.com/davidterranova/userstore/pkg/xerrors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("with invalid query", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name               string
			query              app.GetUserQuery
			expectedErrorClass xerrors.Class
		}{
			{
				name:               "with empty user id",
				query:              app.GetUserQuery{},
				expectedErrorClass: xerrors.ClassBadRequest,
			},
			{
				name: "with invalid user id",
				query: app.GetUserQuery{
					UserId: "not-a-uuid",
				},
				expectedErrorClass: xerrors.ClassBadRequest,
			},
		}

		for i := range tests {
			test := tests[i]
			t.Run(test.name, func(t *testing.T) {
				t.Parallel()

				ctr := setup(t)
				handler := app.NewGetUserHandler(ctr.UserRepository)

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

		ctr := setup(t)
		handler := app.NewGetUserHandler(ctr.UserRepository)
		repoUser := domain.NewUser("john", "doe", "jdoe@userstore.local")

		ctr.UserRepository.EXPECT().
			GetUser(ctx, repoUser.GetId()).
			Times(1).
			Return(repoUser, nil)

		query := app.GetUserQuery{
			UserId: repoUser.GetId().String(),
		}
		user, err := handler.Handle(ctx, query)
		assert.NoError(t, err)
		assert.Equal(t, user, repoUser)
	})

	t.Run("with user not found", func(t *testing.T) {
		t.Parallel()

		ctr := setup(t)
		handler := app.NewGetUserHandler(ctr.UserRepository)
		userId := uuid.New()

		ctr.UserRepository.EXPECT().
			GetUser(ctx, gomock.Any()).
			Times(1).
			Return(nil, adapter.ErrUserNotFound)

		query := app.GetUserQuery{
			UserId: userId.String(),
		}

		user, err := handler.Handle(ctx, query)
		assert.ErrorIs(t, err, adapter.ErrUserNotFound)
		assert.Nil(t, user)

		var classError *xerrors.ClassError
		assert.ErrorAs(t, err, &classError)
		classError, _ = err.(*xerrors.ClassError)
		assert.Equal(t, xerrors.ClassNotFound, classError.Class())
	})

	t.Run("with unexpected repository error", func(t *testing.T) {
		t.Parallel()

		ctr := setup(t)
		handler := app.NewGetUserHandler(ctr.UserRepository)
		userId := uuid.New()

		ctr.UserRepository.EXPECT().
			GetUser(ctx, gomock.Any()).
			Times(1).
			Return(nil, errors.New("unexpected error"))

		query := app.GetUserQuery{
			UserId: userId.String(),
		}

		user, err := handler.Handle(ctx, query)
		assert.Error(t, err)
		assert.Nil(t, user)

		var classError *xerrors.ClassError
		assert.ErrorAs(t, err, &classError)
		classError, _ = err.(*xerrors.ClassError)
		assert.Equal(t, xerrors.ClassInternal, classError.Class())
	})
}
