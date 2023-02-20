package guildedgo

import (
	"errors"
)

type ListItem struct {
	ID                 string `json:"id"`
	ServerID           string `json:"serverId"`
	ChannelID          string `json:"channelId"`
	Messsage           string `json:"message"`
	Mentions           `json:"mentions,omitempty"`
	CreatedAt          string `json:"createdAt"`
	CreatedBy          string `json:"createdBy"`
	CreatedByWebhookID string `json:"createdByWebhookId,omitempty"`
	UpdatedAt          string `json:"updatedAt",omitempty`
	UpdatedBy          string `json:"updatedBy,omitempty"`
	ParentListItemID   string `json:"parentListItemId,omitempty"`
	CompletedAt        string `json:"completedAt,omitempty"`
	CompletedBy        string `json:"completedBy,omitempty"`
	Note               struct {
		CreatedAt string `json:"createdAt"`
		CreatedBy string `json:"createdBy"`
		UpdatedAt string `json:"updatedAt,omitempty"`
		UpdatedBy string `json:"updatedBy"`
		Mentions  `json:"mentions,omitempty"`
		Content   string `json:"content"`
	} `json:"note"`
}

type ListItemSummary struct {
	ID                 string `json:"id"`
	ServerID           string `json:"serverId"`
	ChannelID          string `json:"channelId"`
	Messsage           string `json:"message"`
	Mentions           `json:"mentions,omitempty"`
	CreatedAt          string `json:"createdAt"`
	CreatedBy          string `json:"createdBy"`
	CreatedByWebhookID string `json:"createdByWebhookId,omitempty"`
	UpdatedAt          string `json:"updatedAt,omitempty"`
	UpdatedBy          string `json:"updatedBy,omitempty"`
	ParentListItemID   string `json:"parentListItemId,omitempty"`
	CompletedAt        string `json:"completedAt,omitempty"`
	CompletedBy        string `json:"completedBy,omitempty"`
	Note               struct {
		CreatedAt string `json:"createdAt"`
		CreatedBy string `json:"createdBy"`
		UpdatedAt string `json:"updatedAt,omitempty"`
		UpdatedBy string `json:"updatedBy"`
	} `json:"note,omitempty"`
}

type ListObject struct {
	Message string `json:"message"`
	Note    struct {
		Content string `json:"content"`
	} `json:"note"`
}

type listEndpoints struct{}

func (e *listEndpoints) Default(channelID string) string {
	return guildedApi + "/channels/" + channelID + "/items"
}

func (e *listEndpoints) Get(channelID, listItemID string) string {
	return guildedApi + "/channels/" + channelID + "/items/" + listItemID
}

func (e *listEndpoints) Complete(channelID, listItemID string) string {
	return guildedApi + "/channels/" + channelID + "/items/" + listItemID + "/complete"
}

type ListService interface {
	CreateListItem(channelID string, listObject ListObject) (*ListItem, error)
	GetChannelListItems(channelID string) ([]ListItemSummary, error)
	GetListItem(channelID, listItemID string) (*ListItem, error)
	UpdateListItem(channelID, listItemID string, listObject ListObject) (*ListItem, error)
	DeleteListItem(channelID, listItemID string) error
	CompleteListItem(channelID, listItemID string) error
	UncompleteListItem(channelID, listItemID string) error
}

type listService struct {
	client    *Client
	endpoints *listEndpoints
}

var _ ListService = &listService{
	endpoints: &listEndpoints{},
}

func (service *listService) CreateListItem(channelID string, listObject ListObject) (*ListItem, error) {
	var listItem ListItem
	err := service.client.PostRequestV2(service.endpoints.Default(channelID), &listObject, &listItem)
	if err != nil {
		return nil, errors.New("Error creating list item: " + err.Error())
	}

	return &listItem, nil
}

func (service *listService) GetChannelListItems(channelID string) ([]ListItemSummary, error) {
	var listItems []ListItemSummary
	err := service.client.GetRequestV2(service.endpoints.Default(channelID), &listItems)
	if err != nil {
		return nil, errors.New("Error getting channel list items: " + err.Error())
	}

	return listItems, nil
}

func (service *listService) GetListItem(channelID, listItemID string) (*ListItem, error) {
	var listItem ListItem
	err := service.client.GetRequestV2(service.endpoints.Get(channelID, listItemID), &listItem)
	if err != nil {
		return nil, errors.New("Error getting list item: " + err.Error())
	}

	return &listItem, nil
}

func (service *listService) UpdateListItem(channelID, listItemID string, listObject ListObject) (*ListItem, error) {
	var listItem ListItem
	err := service.client.PutRequestV2(service.endpoints.Get(channelID, listItemID), &listObject, &listItem)
	if err != nil {
		return nil, errors.New("Error updating list item: " + err.Error())
	}

	return &listItem, nil
}

func (service *listService) DeleteListItem(channelID, listItemID string) error {
	_, err := service.client.DeleteRequest(service.endpoints.Get(channelID, listItemID))
	if err != nil {
		return errors.New("Error deleting list item: " + err.Error())
	}

	return nil
}

func (service *listService) CompleteListItem(channelID, listItemID string) error {
	_, err := service.client.PostRequest(service.endpoints.Complete(channelID, listItemID), nil)
	if err != nil {
		return errors.New("Error completing list item: " + err.Error())
	}

	return nil
}

func (service *listService) UncompleteListItem(channelID, listItemID string) error {
	_, err := service.client.DeleteRequest(service.endpoints.Complete(channelID, listItemID))
	if err != nil {
		return errors.New("Error uncompleting list item: " + err.Error())
	}

	return nil
}
