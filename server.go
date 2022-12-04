package guildedgo

import (
	"errors"
	"fmt"

	"github.com/itschip/guildedgo/internal/endpoints"
)

type Server struct {
	ID               string `json:"id"`
	OwnerID          string `json:"ownerid"`
	Type             string `json:"type"`
	Name             string `json:"name"`
	URL              string `json:"url"`
	About            string `json:"about"`
	Avatar           string `json:"avatar"`
	Banner           string `json:"banner"`
	Timezone         string `json:"timezone"`
	IsVerified       bool   `json:"isVerified"`
	DefaultChannelId string `json:"defaultChannelId"`
	CreatedAt        string `json:"createdAt"`
}

type ServerResponse struct {
	Server `json:"server"`
}

type ServerService interface {
	GetServer(serverId string) (*Server, error)
}

type serverService struct {
	client *Client
}

var _ ServerService = &serverService{}

func (ss *serverService) GetServer(serverId string) (*Server, error) {
	endpoint := endpoints.ServerEndpoint(serverId)

	var server ServerResponse
	err := ss.client.GetRequestV2(endpoint, &server)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get server. Error: %v", err.Error()))
	}

	return &server.Server, nil
}
