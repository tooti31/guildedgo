package guildedgo

type CalendarEvent struct {
	ID           int    `json:"id"`
	ServerID     string `json:"serverId"`
	ChannelID    string `json:"channelId"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	Location     string `json:"location,omitempty"`
	URL          string `json:"url,omitempty"`
	Color        int    `json:"color,omitempty"`
	RSVPLimit    int    `json:"rsvpLimit,omitempty"`
	StartsAt     string `json:"startsAt"`
	Duration     string `json:"duration,omitempty"`
	IsPrivate    bool   `json:"isPrivate,omitempty"`
	Mentions     `json:"mentions,omitempty"`
	CreatedAt    string `json:"createdAt"`
	CreatedBy    string `json:"createdBy"`
	Cancellation struct {
		Description string `json:"description,omitempty"`
		CreatedBy   string `json:"createdBy,omitempty"`
	} `json:"cancellation,omitempty"`
}
