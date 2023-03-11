package app

import "github.com/davidterranova/userstore/internal/model"

type App struct {
	GetUser    *GetUserHandler
	CreateUser *CreateUserHandler
}

func New(userRepository model.UserRepository) *App {
	return &App{
		GetUser:    NewGetUserHandler(userRepository),
		CreateUser: NewCreateUserHandler(userRepository),
	}
}
