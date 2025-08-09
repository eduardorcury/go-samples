package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type userService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *userService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) CreateUser(ctx context.Context, user *User) (*User, error) {
	url := fmt.Sprintf("http://user-service/%s/addresses", user.ID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get address from API: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the response: %v", err)
	}

	var addressResponse Address
	if err := json.Unmarshal(body, &addressResponse); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	user = &User{
		ID:      user.ID,
		Email:   user.Email,
		Name:    user.Name,
		Address: addressResponse,
	}

	return s.userRepository.createNewUser(ctx, user)
}
