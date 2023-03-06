package guildedgo

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/itschip/guildedgo/endpoints"
)

type ServerChannel struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	Topic      string `json:"topic"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	ServerID   string `json:"serverId"`
	ParentID   string `json:"parentId"`
	CategoryID string `json:"categoryId"`
	GroupID    string `json:"groupId"`
	IsPublic   bool   `json:"isPublic"`
	ArchivedBy string `json:"archivedBy"`
	ArchivedAt string `json:"archivedAt"`
}

type Mentions struct {
	// Info on mentioned users (min items 1)
	Users MentionsUser `json:"users,omitempty"`

	// Info on mentioned channels (min items 1)
	Channels []MentionsChannel `json:"channels,omitempty"`

	// Info on mentioned roles (min items 1)
	Roles []MentionsRole `json:"roles,omitempty"`

	// If @everyone was mentioned
	Everyone bool `json:"everyone,omitempty"`

	// If @here was mentioned
	Here bool `json:"here,omitempty"`
}

type MentionsRole struct {
	// The ID of the role
	ID int `json:"id"`
}

type MentionsChannel struct {
	// The ID of the channel
	ID string `json:"id"`
}

type MentionsUser struct {
	// The ID of the user
	ID string `json:"id"`
}

type NewChannelObject struct {
	Name       string `json:"name"`
	Topic      string `json:"topic,omitempty"`
	IsPublic   bool   `json:"isPublic,omitempty"`
	Type       string `json:"type"`
	ServerID   string `json:"serverId,omitempty"`
	GroupID    string `json:"groupId,omitempty"`
	CategoryID string `json:"categoryId,omitempty"`
}

type UpdateChannelObject struct {
	Name     string `json:"name,omitempty"`
	Topic    string `json:"topic,omitempty"`
	IsPublic string `json:"isTopic,omitempty"`
}

type ServerChannelResponse struct {
	Channel ServerChannel `json:"channel"`
}

type ChannelService interface {
	CreateChannel(channelObject *NewChannelObject) (*ServerChannel, error)
	GetChannel(channelId string) (*ServerChannel, error)
	UpdateChannel(channelId string, channelObject *UpdateChannelObject) (*ServerChannel, error)
	DeleteChannel(channelId string) error
	SendMessage(channelId string, message *MessageObject) (*ChatMessage, error)
	GetMessages(channelId string, getObject *GetMessagesObject) (*[]ChatMessage, error)
	GetMessage(channelId string, messageId string) (*ChatMessage, error)
	UpdateChannelMessage(channelId string, messageId string, newMessage *MessageObject) (*ChatMessage, error)
	DeleteChannelMessage(channelId string, messageId string) error
}

type channelService struct {
	client *Client
}

var _ ChannelService = &channelService{}

// CreateChannel returns the newly created channel.
// Only server channels are supported at this time (coming soon™: DM Channels!)
func (cs *channelService) CreateChannel(channelObject *NewChannelObject) (*ServerChannel, error) {
	endpoint := endpoints.ChannelEndpoint()

	channelObject.ServerID = cs.client.ServerID

	resp, err := cs.client.PostRequest(endpoint, &channelObject)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to create new channel. Error: \n%v", err.Error()))
	}

	var serverChannel ServerChannelCreated
	err = json.Unmarshal(resp, &serverChannel)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to unmarshal ServerChannel response. Error: \n%v", err.Error()))
	}

	return &serverChannel.Channel, nil
}

// GetChannel returns a channel by channelId.
// Only server channels are supported at this time (coming soon™: DM Channels!)
func (cs *channelService) GetChannel(channelId string) (*ServerChannel, error) {
	endpoint := endpoints.ChannelEndpointWithID(channelId)

	var serverChannel ServerChannelResponse
	err := cs.client.GetRequestV2(endpoint, &serverChannel)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get channel. Error: \n%v", err.Error()))
	}

	return &serverChannel.Channel, nil
}

// UpdateChannel returns the updated channel.
// Only server channels are supported at this time (coming soon™: DM Channels!)
func (cs *channelService) UpdateChannel(channelId string, channelObject *UpdateChannelObject) (*ServerChannel, error) {
	endpoint := endpoints.ChannelEndpointWithID(channelId)

	var serverChannel ServerChannelResponse
	err := cs.client.PatchRequest(endpoint, &channelObject, &serverChannel)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to update  channel. Error: \n%v", err.Error()))
	}

	return &serverChannel.Channel, nil
}

// DeleteChannel does not return anything
// Only server channels are supported at this time (coming soon™: DM Channels!)
func (cs *channelService) DeleteChannel(channelId string) error {
	endpoint := endpoints.ChannelEndpointWithID(channelId)
	_, err := cs.client.DeleteRequest(endpoint)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to delete channel. Error: \n%v", err.Error()))
	}

	return nil
}

func (cs *channelService) SendMessage(channelId string, message *MessageObject) (*ChatMessage, error) {
	endpoint := endpoints.CreateMessageEndpoint(channelId)

	resp, err := cs.client.PostRequest(endpoint, &message)
	if err != nil {
		return nil, err
	}

	var msg MessageResponse
	err = json.Unmarshal(resp, &msg)
	if err != nil {
		return nil, err
	}

	return &msg.Message, err
}

func (cs *channelService) UpdateChannelMessage(channelId string, messageId string, newMessage *MessageObject) (*ChatMessage, error) {
	endpoint := endpoints.UpdateChannelMessageEndpoint(channelId, messageId)

	resp, err := cs.client.PutRequest(endpoint, &newMessage)
	if err != nil {
		return nil, err
	}

	var msg MessageResponse
	err = json.Unmarshal(resp, &msg)
	if err != nil {
		return nil, err
	}

	return &msg.Message, err
}

// GetMessages TODO: add support for params
func (cs *channelService) GetMessages(channelId string, getObject *GetMessagesObject) (*[]ChatMessage, error) {
	endpoint := endpoints.GetChannelMessagesEndpoint(channelId)

	resp, err := cs.client.GetRequest(endpoint)
	if err != nil {
		return nil, err
	}

	// Abstract this functionality in GetRequest, as for the rest below and above
	var msgs AllMessagesResponse
	err = json.Unmarshal(resp, &msgs)
	if err != nil {
		return nil, err
	}

	return &msgs.Messages, nil
}

// GetMessage Get a message from a channel
func (cs *channelService) GetMessage(channelId string, messageId string) (*ChatMessage, error) {
	endpoint := endpoints.GetChannelMessageEndpoint(channelId, messageId)

	resp, err := cs.client.GetRequest(endpoint)
	if err != nil {
		return nil, err
	}

	var msg MessageResponse
	err = json.Unmarshal(resp, &msg)
	if err != nil {
		return nil, err
	}

	return &msg.Message, nil
}

func (cs *channelService) DeleteChannelMessage(channelId string, messageId string) error {
	endpoint := endpoints.GetChannelMessageEndpoint(channelId, messageId)

	_, err := cs.client.DeleteRequest(endpoint)
	if err != nil {
		return err
	}

	return nil
}
