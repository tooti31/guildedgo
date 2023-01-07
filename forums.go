package guildedgo

import (
	"errors"
	"fmt"

	"github.com/itschip/guildedgo/endpoints"
)

type ForumTopic struct {
	// The ID of the forum topic
	ID string `json:"id"`

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

type ForumService interface{}

type forumService struct {
	client *Client
}

var _ ForumService = &forumService{}

func (f *forumService) CreateForumTopic(channelId string, forumTopicObject *ForumTopicObject) (*ForumTopic, error) {
	endpoint := endpoints.ForumTopicEndpoint(channelId)

	var forumTopic ForumTopic
	err := f.client.PostRequestV2(endpoint, &forumTopicObject, &forumTopic)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to create new forum topic. Error: \n%v", err.Error()))
	}

	return &forumTopic, nil
}

func (f *forumService) GetForumTopics(channelId string) (*[]ForumTopicSummary, error) {
	endpoint := endpoints.ForumTopicEndpoint(channelId)

	var forumTopicSummary []ForumTopicSummary

	err := f.client.GetRequestV2(endpoint, &forumTopicSummary)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get forum topics. Error: \n%v", err.Error()))
	}

	return &forumTopicSummary, nil
}

func (f *forumService) GetForumTopic(channelId string, forumTopicId int) (*ForumTopic, error) {
	endpoint := endpoints.GetForumTopicEndpoint(channelId, forumTopicId)

	var forumTopic ForumTopic

	err := f.client.GetRequestV2(endpoint, &forumTopic)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get forum topic. Error: \n%v", err.Error()))
	}

	return &forumTopic, nil
}

func (f *forumService) UpdateForumTopic(channelId string, forumTopicId int) {
}

func (f *forumService) DeleteForumTopic() {
}

func (f *forumService) PinForumTopic() {
}

func (f *forumService) UnpinForumTopic() {
}

func (f *forumService) LockForumTopic() {
}

func (f *forumService) UnlockForumTopic() {
}
