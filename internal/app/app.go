package app

import "github.com/davidterranova/userstore/internal/model"

type App struct {
	GetUser GetUserHandler
}

func New(userRepository model.UserRepository) *App {
	return &App{
		GetUser: *NewGetUserHandler(userRepository),
	}
}
