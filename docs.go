package guildedgo

import (
	"errors"
	"net/url"
	"strconv"
)

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

type DocObject struct {
	title   string `json:"title"`
	content string `json:"content"`
}

type DocResponse struct {
	Doc `json:"doc"`
}

type DocsResponse struct {
	Docs []Doc `json:"docs"`
}

type docsEndpoints struct{}

func (e *docsEndpoints) Default(channelId string) string {
	return guildedApi + "/channels/" + channelId + "/docs"
}

func (e *docsEndpoints) Get(channelId string, docId int) string {
	return guildedApi + "/channels/" + channelId + "/docs/" + strconv.Itoa(docId)
}

type docsService struct {
	client    *Client
	endpoints *docsEndpoints
}

type DocsService interface {
	Create(channelId string) (*Doc, error)
}

var _ DocsService = &docsService{}

func (service *docsService) Create(channelId string) (*Doc, error) {
	var docResponse DocResponse
	endpoint := service.endpoints.Default(channelId)

	err := service.client.PostRequestV2("POST", endpoint, nil)
	if err != nil {
		return nil, errors.New("Error creating doc. Error: " + err.Error())
	}

	return &docResponse.Doc, nil
}

func (service *docsService) GetDocs(channelId string, doc *DocObject) ([]Doc, error) {
	var docsResponse DocsResponse

	url, err := url.Parse(service.endpoints.Default(channelId))
	if err != nil {
		return nil, errors.New("Error parsing URL. Error: " + err.Error())
	}

	query := url.Query()
	if doc.title != "" {
		query.Add("title", doc.title)
	}
	if doc.content != "" {
		query.Add("content", doc.content)
	}

	url.RawQuery = query.Encode()

	err = service.client.GetRequestV2(url.String(), &docsResponse)
	if err != nil {
		return nil, errors.New("Error getting docs. Error: " + err.Error())
	}

	return docsResponse.Docs, nil
}

func (service *docsService) GetDoc(channelId string, docId int) (*Doc, error) {
	var docResponse DocResponse

	endpoint := service.endpoints.Get(channelId, docId)

	err := service.client.GetRequestV2(endpoint, &docResponse)
	if err != nil {
		return nil, errors.New("Error getting doc. Error: " + err.Error())
	}

	return &docResponse.Doc, nil
}

func (service *docsService) UpdateDoc(channelId string, docId int, doc *DocObject) (*Doc, error) {
	var docResponse DocResponse

	endpoint := service.endpoints.Get(channelId, docId)

	err := service.client.PutRequestV2(endpoint, &doc, &docResponse)
	if err != nil {
		return nil, errors.New("Error updating doc. Error: " + err.Error())
	}

	return &docResponse.Doc, nil
}

func (service *docsService) DeleteDoc(channelId string, docId int) error {
	endpoint := service.endpoints.Get(channelId, docId)

	_, err := service.client.DeleteRequest(endpoint)
	if err != nil {
		return errors.New("Error deleting doc. Error: " + err.Error())
	}

	return nil
}
