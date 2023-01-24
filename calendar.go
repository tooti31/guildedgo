package guildedgo

import (
	"errors"
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

type CalenderEventObject struct {
	Name        string `json:"name,omitempty"`
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

type CalendarEventRsvpResponse struct {
	Rsvp CalendarEventRsvp `json:"calendarEventRsvp"`
}

type CalendarEventRsvpsResponse struct {
	Rsvps []CalendarEventRsvp `json:"calendarEventRsvps"`
}

type calenderEndpoints struct{}

func (e *calenderEndpoints) Default(channelId string) string {
	return guildedApi + "/channels/" + channelId + "/events"
}

func (e *calenderEndpoints) Get(channelId string, eventId int) string {
	return guildedApi + "/channels/" + channelId + "/events/" + strconv.Itoa(eventId)
}

func (e *calenderEndpoints) RsvpDefault(channelId string, eventId int, userId string) string {
	return guildedApi + "/channels/" + channelId + "/events/" + strconv.Itoa(eventId) + "/rsvps/" + userId
}

func (e *calenderEndpoints) Rsvps(channelId string, eventId int) string {
	return guildedApi + "/channels/" + channelId + "/events/" + strconv.Itoa(eventId) + "/rsvps"
}

type calendarService struct {
	client    *Client
	endpoints *calenderEndpoints
}

type CalendarService interface {
	CreateEvent(channelId string, event *CalenderEventObject) (*CalendarEvent, error)
	GetEvents(channelId string, options *GetEventsOptions) ([]CalendarEvent, error)
	GetEvent(channelId string, eventId int) (*CalendarEvent, error)
	UpdateEvent(channelId string, eventId int, event *CalenderEventObject) (*CalendarEvent, error)
	DeleteEvent(channelId string, eventId int) error
	GetEventRSVP(channelId string, eventId int, userId string) (*CalendarEventRsvp, error)
	CreateOrUpdateEventRSVP(channelId string, eventId int, userId string) (*CalendarEventRsvp, error)
	DeleteEventRSVP(channelId string, eventId int, userId string) error
	GetEventRSVPs(channelId string, eventId int) ([]CalendarEventRsvp, error)
}

var _ CalendarService = &calendarService{
	endpoints: &calenderEndpoints{},
}

func (service *calendarService) CreateEvent(channelId string, event *CalenderEventObject) (*CalendarEvent, error) {
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

func (service *calendarService) UpdateEvent(channelId string, eventId int, event *CalenderEventObject) (*CalendarEvent, error) {
	var calendarEvent CalendarEventResponse

	err := service.client.PatchRequest(service.endpoints.Get(channelId, eventId), event, &calendarEvent)
	if err != nil {
		return nil, errors.New("Failed to update calendar event: " + err.Error())
	}

	return &calendarEvent.Event, nil
}

func (service *calendarService) DeleteEvent(channelId string, eventId int) error {
	_, err := service.client.DeleteRequest(service.endpoints.Get(channelId, eventId))
	if err != nil {
		return errors.New("Failed to delete calendar event: " + err.Error())
	}

	return nil
}

func (service *calendarService) GetEventRSVP(channelId string, eventId int, userId string) (*CalendarEventRsvp, error) {
	var calendarEventRsvp CalendarEventRsvpResponse

	err := service.client.GetRequestV2(service.endpoints.RsvpDefault(channelId, eventId, userId), &calendarEventRsvp)
	if err != nil {
		return nil, errors.New("Failed to get calendar event rsvp: " + err.Error())
	}

	return &calendarEventRsvp.Rsvp, nil
}

func (service *calendarService) CreateOrUpdateEventRSVP(channelId string, eventId int, userId string) (*CalendarEventRsvp, error) {
	var calendarEventRsvp CalendarEventRsvpResponse

	err := service.client.GetRequestV2(service.endpoints.RsvpDefault(channelId, eventId, userId), &calendarEventRsvp)
	if err != nil {
		return nil, errors.New("Failed to create or update calendar event RSVP. Error: " + err.Error())
	}

	return &calendarEventRsvp.Rsvp, nil
}

func (service *calendarService) DeleteEventRSVP(channelId string, eventId int, userId string) error {
	_, err := service.client.DeleteRequest(service.endpoints.RsvpDefault(channelId, eventId, userId))
	if err != nil {
		return errors.New("Failed to delete calendar event RSVP. Error: " + err.Error())
	}

	return nil
}

func (service *calendarService) GetEventRSVPs(channelId string, eventId int) ([]CalendarEventRsvp, error) {
	var calendarEventRsvps CalendarEventRsvpsResponse

	err := service.client.GetRequestV2(service.endpoints.Rsvps(channelId, eventId), &calendarEventRsvps)
	if err != nil {
		return nil, errors.New("Failed to get calendar event RSVPs. Error: " + err.Error())
	}

	return calendarEventRsvps.Rsvps, nil
}
