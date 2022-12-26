package guildedgo

type Doc struct {
	ID        int    `json:"id"`
	ServerID  string `json:"serverId"`
	ChannelID string `json:"channelId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Mentions  `json:"mentions,omitempty"`
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	UpdatedBy string `json:"updatedBy,omitempty"`
}
