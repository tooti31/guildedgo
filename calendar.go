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

type CalendarEventRsvp struct {
	CalendarEventID int    `json:"calendarEventId"`
	ChannelID       string `json:"channelId"`
	ServerID        string `json:"serverId"`
	UserID          string `json:"userId"`
	Status          string `json:"status"`
	CreatedBy       string `json:"createdAt"`
	CreatedAt       string `json:"createdAt"`
	UpdatedBy       string `json:"updatedAt"`
	UpdatedAt       string `json:"updatedAt"`
}
