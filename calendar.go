package guildedgo

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

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
	Duration     int    `json:"duration,omitempty"`
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
	CreatedBy       string `json:"createdBy"`
	CreatedAt       string `json:"createdAt"`
	UpdatedBy       string `json:"updatedBy"`
	UpdatedAt       string `json:"updatedAt"`
}

type CreateCalenderEventObject struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Location    string `json:"location,omitempty"`
	StartsAt    string `json:"startsAt,omitempty"`
	URL         string `json:"url,omitempty"`
	Color       int    `json:"color,omitempty"`
	RSVPLimit   int    `json:"rsvpLimit,omitempty"`
	Duration    int    `json:"duration,omitempty"`
	IsPrivate   bool   `json:"isPrivate,omitempty"`
}

type CalendarEventResponse struct {
	Event CalendarEvent `json:"calendarEvent"`
}

type CalendarEventsResponse struct {
	Events []CalendarEvent `json:"calendarEvents"`
}

type calenderEndpoints struct{}

func (e *calenderEndpoints) Default(channelId string) string {
	return guildedApi + "/channels/" + channelId + "/events"
}

func (e *calenderEndpoints) Get(channelId string, eventId int) string {
	return guildedApi + "/channels/" + channelId + "/events/" + strconv.Itoa(eventId)
}

type calendarService struct {
	client    *Client
	endpoints *calenderEndpoints
}

type CalendarService interface {
	CreateEvent(channelId string, event *CreateCalenderEventObject) (*CalendarEvent, error)
	GetEvents(channelId string, options *GetEventsOptions) ([]CalendarEvent, error)
	GetEvent(channelId string, eventId int) (*CalendarEvent, error)
}

var _ CalendarService = &calendarService{
	endpoints: &calenderEndpoints{},
}

func (service *calendarService) CreateEvent(channelId string, event *CreateCalenderEventObject) (*CalendarEvent, error) {
	var calendarEvent CalendarEventResponse

	err := service.client.PostRequestV2(service.endpoints.Default(channelId), event, &calendarEvent)
	if err != nil {
		return nil, errors.New("Failed to create calendar event: " + err.Error())
	}

	return &calendarEvent.Event, nil
}

type GetEventsOptions struct {
	Before string
	After  string
	Limit  int
}

func (service *calendarService) GetEvents(channelId string, options *GetEventsOptions) ([]CalendarEvent, error) {
	var calendarEvents CalendarEventsResponse

	url, err := url.Parse(service.endpoints.Default(channelId))
	if err != nil {
		return nil, err
	}

	values := url.Query()
	if options.Before != "" {
		values.Add("before", options.Before)
	}
	if options.After != "" {
		values.Add("after", options.After)
	}
	if options.Limit != 0 {
		values.Add("limit", strconv.Itoa(options.Limit))
	}

	url.RawQuery = values.Encode()

	err = service.client.GetRequestV2(url.String(), &calendarEvents)
	if err != nil {
		return nil, errors.New("Failed to get calendar events: " + err.Error())
	}

	return calendarEvents.Events, nil
}

func (service *calendarService) GetEvent(channelId string, eventId int) (*CalendarEvent, error) {
	var calendarEvent CalendarEventResponse

	err := service.client.GetRequestV2(service.endpoints.Get(channelId, eventId), &calendarEvent)
	if err != nil {
		return nil, errors.New("Failed to get calendar event: " + err.Error())
	}

	return &calendarEvent.Event, nil
}
