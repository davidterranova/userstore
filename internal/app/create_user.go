package app

import (
	"context"
	"errors"
	"strings"

	"github.com/davidterranova/userstore/internal/adapter"
	"github.com/davidterranova/userstore/internal/domain"
	"github.com/davidterranova/userstore/pkg/xerrors"
	"github.com/go-playground/validator"
)

type CreateUserCmd struct {
	FirstName string `validate:"required,min=1,max=255"`
	LastName  string `validate:"required,min=1,max=255"`
	Email     string `validate:"required,min=1,max=255,email"`
}

type CreateUserHandler struct {
	validator      *validator.Validate
	userRepository domain.UserRepository
}

func NewCreateUserHandler(userRepository domain.UserRepository) *CreateUserHandler {
	return &CreateUserHandler{
		validator:      validator.New(),
		userRepository: userRepository,
	}
}

func (h *CreateUserHandler) Handle(ctx context.Context, cmd CreateUserCmd) (*domain.User, error) {
	err := h.validator.Struct(cmd)
	if err != nil {
		return nil, xerrors.NewClassError(xerrors.ClassBadRequest, err)
	}

	newUser := domain.NewUser(
		sanitize(cmd.FirstName),
		sanitize(cmd.LastName),
		sanitize(cmd.Email),
	)

	user, err := h.userRepository.CreateUser(ctx, newUser)
	if err != nil {
		if errors.Is(err, adapter.ErrUserAlreadyExist) {
			return nil, xerrors.NewClassError(xerrors.ClassConflict, err)
		}
		return nil, xerrors.NewClassError(xerrors.ClassInternal, err)
	}

	return user, nil
}

func sanitize(input string) string {
	return strings.ToLower(
		strings.TrimSpace(input),
	)
}
