package guildedgo

import (
	"errors"
	"fmt"
	"log"

	"github.com/itschip/guildedgo/endpoints"
)

type ForumTopic struct {
	// The ID of the forum topic
	ID int `json:"id"`

	// The ID of the server
	ServerID string `json:"serverId"`

	// The ID of the channel
	ChannelID string `json:"channelId"`

	// The title of the forum topic (min length 1; max length 500)
	Title string `json:"title"`

	// The ISO 8601 timestamp that the forum topic was created at
	CreatedAt string `json:"createdAt"`

	// The ID of the user who created this forum topic
	// (Note: If this event has createdByWebhookId present, this field will still be populated,
	// but can be ignored. In this case, the value of this field will always be Ann6LewA)
	CreatedBy string `json:"createdBy"`

	// The ID of the webhook who created this forum topic, if it was created by a webhook
	CreatedByWebhookID string `json:"createdByWebhookId,omitempty"`

	// The ISO 8601 timestamp that the forum topic was updated at, if relevant
	UpdatedAt string `json:"updatedAt,omitempty"`

	// The ISO 8601 timestamp that the forum topic was bumped at.
	// This timestamp is updated whenever there is any activity on the posts within the forum topic.
	BumpedAt string `json:"bumpedAt,omitempty"`

	// (default false)
	IsPinned bool `json:"isPinned,omitempty"`

	// (default false)
	IsLocked bool `json:"isLocked,omitempty"`

	// The content of the forum topic
	Content string `json:"content"`

	Mentions `json:"mentions,omitempty"`
}

type ForumTopicSummary struct {
	// The ID of the forum topic
	ID int `json:"id"`

	// The ID of the server
	ServerID string `json:"serverId"`

	// The ID of the channel
	ChannelID string `json:"channelId"`

	// The title of the forum topic (min length 1; max length 500)
	Title string `json:"title"`

	// The ISO 8601 timestamp that the forum topic was created at
	CreatedAt string `json:"createdAt"`

	// The ID of the user who created this forum topic
	// (Note: If this event has createdByWebhookId present, this field will still be populated,
	// but can be ignored. In this case, the value of this field will always be Ann6LewA)
	CreatedBy string `json:"createdBy"`

	// The ID of the webhook who created this forum topic, if it was created by a webhook
	CreatedByWebhookID string `json:"createdByWebhookId,omitempty"`

	// The ISO 8601 timestamp that the forum topic was updated at, if relevant
	UpdatedAt string `json:"updatedAt,omitempty"`

	// The ISO 8601 timestamp that the forum topic was bumped at.
	// This timestamp is updated whenever there is any activity on the posts within the forum topic.
	BumpedAt string `json:"bumpedAt,omitempty"`

	// (default false)
	IsPinned bool `json:"isPinned,omitempty"`

	// (default false)
	IsLocked bool `json:"isLocked,omitempty"`
}

type ForumTopicComment struct {
	ID           int    `json:"id"`
	Content      string `json:"content"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
	ChannelID    string `json:"channelId"`
	ForumTopicID int    `json:"forumTopicId"`
	CreatedBy    string `json:"createdBy"`
	Mentions     `json:"mentions"`
}

type ForumTopicObject struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateTopicObject struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type forumEndpoints struct{}

func (e *forumEndpoints) Default(channelId string) string {
	return guildedApi + "/channels/" + channelId + "/topics"
}

func (e *forumEndpoints) Get(channelId string, forumTopicId int) string {
	return guildedApi + "/channels/" + channelId + "/topics/" + fmt.Sprint(forumTopicId)
}

func (e *forumEndpoints) Pin(channelId string, forumTopicId int) string {
	return guildedApi + "/channels/" + channelId + "/topics/" + fmt.Sprint(forumTopicId) + "/pin"
}

func (e *forumEndpoints) Lock(channelId string, forumTopicId int) string {
	return guildedApi + "/channels/" + channelId + "/topics/" + fmt.Sprint(forumTopicId) + "/lock"
}

type forumService struct {
	client    *Client
	endpoints *forumEndpoints
}

type ForumService interface {
	CreateForumTopic(channelId string, forumTopicObject *ForumTopicObject) (*ForumTopic, error)
	GetForumTopics(channelId string) (*[]ForumTopicSummary, error)
	GetForumTopic(channelId string, forumTopicId int) (*ForumTopic, error)
	UpdateForumTopic(channelId string, forumTopicId int, updateTopicObject *UpdateTopicObject) (*ForumTopic, error)
	DeleteForumTopic(channelId string, forumTopicId int) error
	PinForumTopic(channelId string, forumTopicId int) error
	UnpinForumTopic(channelId string, forumTopicId int) error
	LockForumTopic(channelId string, forumTopicId int) error
	UnlockForumTopic(channelId string, forumTopicId int) error
}

var _ ForumService = &forumService{
	endpoints: &forumEndpoints{},
}

func (service *forumService) CreateForumTopic(channelId string, forumTopicObject *ForumTopicObject) (*ForumTopic, error) {
	endpoint := service.endpoints.Default(channelId)

	var forumTopic ForumTopic
	err := service.client.PostRequestV2(endpoint, &forumTopicObject, &forumTopic)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to create new forum topic. Error: \n%v", err.Error()))
	}

	return &forumTopic, nil
}

func (service *forumService) GetForumTopics(channelId string) (*[]ForumTopicSummary, error) {
	endpoint := service.endpoints.Default(channelId)

	var forumTopicSummary []ForumTopicSummary

	err := service.client.GetRequestV2(endpoint, &forumTopicSummary)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get forum topics. Error: \n%v", err.Error()))
	}

	return &forumTopicSummary, nil
}

func (service *forumService) GetForumTopic(channelId string, forumTopicId int) (*ForumTopic, error) {
	endpoint := service.endpoints.Get(channelId, forumTopicId)

	var forumTopic ForumTopic

	err := service.client.GetRequestV2(endpoint, &forumTopic)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get forum topic. Error: \n%v", err.Error()))
	}

	return &forumTopic, nil
}

func (service *forumService) UpdateForumTopic(channelId string, forumTopicId int, topicObject *UpdateTopicObject) (*ForumTopic, error) {
	endpoint := service.endpoints.Get(channelId, forumTopicId)

	var forumTopic ForumTopic

	err := service.client.PatchRequest(endpoint, &topicObject, &forumTopic)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to update forum topic. Error: \n%v", err.Error()))
	}

	return &forumTopic, nil
}

func (service *forumService) DeleteForumTopic(channelId string, forumTopicId int) error {
	endpoint := service.endpoints.Get(channelId, forumTopicId)

	_, err := service.client.DeleteRequest(endpoint)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to delete forum topic. Error: \n%v", err.Error()))
	}

	return nil
}

func (service *forumService) PinForumTopic(channelId string, forumTopicId int) error {
	endpoint := service.endpoints.Pin(channelId, forumTopicId)

	_, err := service.client.PutRequest(endpoint, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to pin forum topic. Error: \n%v", err.Error()))
	}

	return nil
}

func (service *forumService) UnpinForumTopic(channelId string, forumTopicId int) error {
	endpoint := service.endpoints.Pin(channelId, forumTopicId)

	_, err := service.client.DeleteRequest(endpoint)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to unpin forum topic. Error: \n%v", err.Error()))
	}

	return nil
}

func (service *forumService) LockForumTopic(channelId string, forumTopicId int) error {
	endpoint := service.endpoints.Lock(channelId, forumTopicId)

	_, err := service.client.PutRequest(endpoint, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to lock forum topic. Error: \n%v", err.Error()))
	}

	return nil
}

func (service *forumService) UnlockForumTopic(channelId string, forumTopicId int) error {
	endpoint := service.endpoints.Lock(channelId, forumTopicId)

	_, err := service.client.DeleteRequest(endpoint)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to unlock forum topic. Error: \n%v", err.Error()))
	}

	return nil
}

func (service *forumService) CreateTopicComment() {
}
