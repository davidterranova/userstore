package app

import "github.com/davidterranova/userstore/internal/domain"

type App struct {
	GetUser    *GetUserHandler
	CreateUser *CreateUserHandler
}

func New(userRepository domain.UserRepository) *App {
	return &App{
		GetUser:    NewGetUserHandler(userRepository),
		CreateUser: NewCreateUserHandler(userRepository),
	}
}
