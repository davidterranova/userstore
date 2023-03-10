package app

import (
	"context"
	"errors"

	"github.com/davidterranova/userstore/internal/adapter"
	"github.com/davidterranova/userstore/internal/model"
	"github.com/davidterranova/userstore/pkg/xerrors"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type GetUserQuery struct {
	UserId string `validate:"required"`
}

type GetUserHandler struct {
	validator      *validator.Validate
	userRepository model.UserRepository
}

func NewGetUserHandler(userRepository model.UserRepository) *GetUserHandler {
	return &GetUserHandler{
		validator:      validator.New(),
		userRepository: userRepository,
	}
}

func (h *GetUserHandler) Handle(ctx context.Context, query GetUserQuery) (*model.User, error) {
	err := h.validator.Struct(query)
	if err != nil {
		return nil, xerrors.NewClassError(xerrors.ClassBadRequest, err)
	}

	userId, err := uuid.Parse(query.UserId)
	if err != nil {
		return nil, xerrors.NewClassError(xerrors.ClassBadRequest, err)
	}

	user, err := h.userRepository.GetUser(ctx, userId)
	if err != nil {
		if errors.Is(err, adapter.ErrUserNotFound) {
			return nil, xerrors.NewClassError(xerrors.ClassNotFound, err)
		}
		return nil, xerrors.NewClassError(xerrors.ClassInternal, err)
	}

	return user, nil
}
