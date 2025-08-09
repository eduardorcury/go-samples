package http

import "context"

type User struct {
	ID      string
	Email   string
	Name    string
	Address Address
}

type Address struct {
	Street string
	City   string
}

type UserRepository interface {
	createNewUser(ctx context.Context, user *User) (*User, error)
}

type UserService interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
}
