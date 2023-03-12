package app_test

import (
	"testing"

	"github.com/davidterranova/userstore/internal/domain"
	"github.com/golang/mock/gomock"
)

type TestContainer struct {
	UserRepository *domain.MockUserRepository
}

func setup(t *testing.T) *TestContainer {
	t.Helper()

	ctrl := gomock.NewController(t)

	return &TestContainer{
		UserRepository: domain.NewMockUserRepository(ctrl),
	}
}
