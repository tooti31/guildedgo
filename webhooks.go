package guildedgo

import "errors"

type WebhookObject struct {
	Name      string `json:"name"`
	ChannelID string `json:"channelId"`
}

type webhookEndpoints struct{}

func (e *webhookEndpoints) Default(serverId string) string {
	return guildedApi + "/servers/" + serverId + "/webhooks"
}

func (e *webhookEndpoints) Get(serverId string, webhookId string) string {
	return guildedApi + "/servers/" + serverId + "/webhooks/" + webhookId
}

type WebhookService interface {
	CreateWebhook(serverId string, webhookObject *WebhookObject) (*Webhook, error)
	GetWebhooks(serverId string) ([]Webhook, error)
	GetWebhook(serverId string, webhookId string) (*Webhook, error)
	UpdateWebhook(serverId string, webhookId string, webhookObject *WebhookObject) (*Webhook, error)
	DeleteWebhook(serverId string, webhookId string) error
}

type webhookService struct {
	endpoints *webhookEndpoints
	client    *Client
}

var _ WebhookService = &webhookService{
	endpoints: &webhookEndpoints{},
}

func (service *webhookService) CreateWebhook(serverId string, webhookObject *WebhookObject) (*Webhook, error) {
	var webhook Webhook

	err := service.client.PostRequestV2(service.endpoints.Default(serverId), &webhookObject, &webhook)
	if err != nil {
		return nil, errors.New("Failed to create webhook: " + err.Error())
	}

	return &webhook, nil
}

func (service *webhookService) GetWebhooks(serverId string) ([]Webhook, error) {
	var webhooks []Webhook

	err := service.client.GetRequestV2(service.endpoints.Default(serverId), &webhooks)
	if err != nil {
		return nil, errors.New("Failed to get webhooks: " + err.Error())
	}

	return webhooks, nil
}

func (service *webhookService) GetWebhook(serverId string, webhookId string) (*Webhook, error) {
	var webhook Webhook

	err := service.client.GetRequestV2(service.endpoints.Get(serverId, webhookId), &webhook)
	if err != nil {
		return nil, errors.New("Failed to get webhook: " + err.Error())
	}

	return &webhook, nil
}

func (service *webhookService) UpdateWebhook(serverId string, webhookId string, webhookObject *WebhookObject) (*Webhook, error) {
	var webhook Webhook

	err := service.client.PutRequestV2(service.endpoints.Get(serverId, webhookId), &webhookObject, &webhook)
	if err != nil {
		return nil, errors.New("Failed to update webhook: " + err.Error())
	}

	return &webhook, nil
}

func (service *webhookService) DeleteWebhook(serverId string, webhookId string) error {
	_, err := service.client.DeleteRequest(service.endpoints.Get(serverId, webhookId))
	if err != nil {
		return errors.New("Failed to delete webhook: " + err.Error())
	}

	return nil
}
