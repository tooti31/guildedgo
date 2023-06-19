package guildedgo

import "errors"

type UserResponse struct {
	User `json:"user"`
}

type UserService interface {
	GetOwnUser() (*UserResponse, error)
}

type userEndpoint struct{}

func (e *userEndpoint) Default() string {
	return guildedApi + "/users/@me"
}

type userService struct {
	client    *Client
	endpoints *userEndpoint
}

func (s *userService) GetOwnUser() (*UserResponse, error) {
	var user UserResponse
	err := s.client.GetRequestV2(s.endpoints.Default(), &user)
	if err != nil {
		return nil, errors.New("failed to get social links: " + err.Error())
	}
	return &user, nil
}
