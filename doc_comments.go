package guildedgo

type DocComment struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	ChannelID string `json:"channelId"`
	DocID     int    `json:"docId"`
	Mentions  `json:"mentions,omitempty"`
}
