package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	id uuid.UUID

	createdAt time.Time

	firstName string
	lastName  string
	email     string
}

func NewUser(firstName, lastName, email string) User {
	return User{
		id:        uuid.New(),
		createdAt: time.Now().UTC().Round(time.Millisecond),

		firstName: firstName,
		lastName:  lastName,
		email:     email,
	}
}

func (u User) GetId() uuid.UUID {
	return u.id
}

func (u User) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u User) GetFirstName() string {
	return u.firstName
}

func (u User) GetLastName() string {
	return u.lastName
}

func (u User) GetEmail() string {
	return u.email
}
