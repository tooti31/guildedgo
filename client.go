package guildedgo

import "net/http"

type Client struct {
	Token    string
	ServerID string
	client   *http.Client

	Channel ChannelService
	Members MembersService
	Roles   RoleService
	Server  ServerService
	Events  map[string]Event
}

type Event struct {
	Callback func(*Client, any)
	Type     *interface{}
}

type Config struct {
	Token    string
	ServerID string
}

func NewClient(config *Config) *Client {
	c := &Client{
		Token:    config.Token,
		ServerID: config.ServerID,
		client:   http.DefaultClient,
	}

	c.Channel = &channelService{client: c}
	c.Members = &membersService{client: c}
	c.Roles = &roleService{client: c}
	c.Server = &serverService{client: c}

	c.Events = make(map[string]Event)

	return c
}
