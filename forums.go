package guildedgo

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
