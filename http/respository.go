package http

import "context"

type userRepository struct {
	users map[string]*User
}

func NewUserRepository() *userRepository {
	return &userRepository{
		users: make(map[string]*User),
	}
}

func (r *userRepository) createNewUser(ctx context.Context, user *User) (*User, error) {
	r.users[user.ID] = user
	return user, nil
}
