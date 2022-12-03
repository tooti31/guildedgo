package guildedgo

type ChatMessage struct {
	// The ID of the message
	ID string `json:"id"`

	// The type of chat message. "system" messages are generated by Guilded,
	// while "defualt" messages are user or bot generated.
	Type string `json:"type"`

	// The ID for th server
	ServerID string `json:"serverId"`

	// The ID of the channel
	ChannelID string `json:"channelId"`

	// The content of the message
	Content string `json:"content"`

	// Message IDs that were replied to
	ReplyMessageIds []string `json:"replyMessageIds"`

	// If set, this message will only be seen by those mentioned or replied to.
	IsPrivate bool `json:"isPrivate"`

	// The ISO 8601 timestamp that the message was created at.
	CreatedAt string `json:"createdAt"`

	// The ID for the user who created this messsage
	// (Note: if this event has `createdByWebhookId` present,
	// this field will still be populated, but can be ignored.
	// In this case, the value of this field will always be Ann6LewA)
	CreatedBy string `json:"createdBy"`

	// The ID of the webhook who created this message, if was created by a webhook
	CreatedByWebhookId string `json:"createdByWebhookId"`

	// The IOSO 8601 timestamp that the message was updated at, if relevant
	UpdatedAt string `json:"updatedAt"`
}

type ChannelMessage struct {
	Id        string `json:"id"`
	Type      string `json:"type"`
	ServerId  string `json:"serverId"`
	ChannelId string `json:"channelId"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
}

type MessageObject struct {
	Content         string `json:"content,omitempty"`
	IsPrivate       string `json:"isPrivate,omitempty"`
	IsSilent        string `json:"isSilent,omitempty"`
	ReplyMessageIds string `json:"replyMessageIds,omitempty"`
}

type MessageResponse struct {
	Message ChatMessage `json:"message"`
}

type GetMessageResponse struct {
	Message ChatMessage `json:"message"`
}

type GetMessagesObject struct {
	Before         string `json:"before"`
	After          string `json:"after"`
	Limit          int    `json:"limit"`
	IncludePrivate bool   `json:"includePrivate"`
}

type AllMessagesResponse struct {
	Messages []ChatMessage `json:"messages"`
}
