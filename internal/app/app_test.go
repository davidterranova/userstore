package app_test

import (
	"testing"

	"github.com/davidterranova/userstore/internal/model"
	"github.com/golang/mock/gomock"
)

type TestContainer struct {
	UserRepository *model.MockUserRepository
}

func setup(t *testing.T) *TestContainer {
	t.Helper()

	ctrl := gomock.NewController(t)

	return &TestContainer{
		UserRepository: model.NewMockUserRepository(ctrl),
	}
}
