package guildedgo

import (
	"errors"
	"strconv"
)

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

type docCommentEndpoints struct{}

func (e *docCommentEndpoints) Default(channelID string, docID int) string {
	return guildedApi + "/channels/" + channelID + "/docs/" + strconv.Itoa(docID) + "/comments"
}

type DocCommentService interface {
	Create(channelID string, docID int, content string) (*DocComment, error)
	GetComments(channelID string, docID int) ([]DocComment, error)
}

type docCommentService struct {
	client    *Client
	endpoints *docCommentEndpoints
}

var _ DocCommentService = &docCommentService{}

func (s *docCommentService) Create(channelID string, docID int, content string) (*DocComment, error) {
	var comment DocComment
	err := s.client.PostRequestV2(s.endpoints.Default(channelID, docID), content, &comment)
	if err != nil {
		return nil, errors.New("error creating doc comment: " + err.Error())
	}

	return &comment, nil
}

func (s *docCommentService) GetComments(channelID string, docID int) ([]DocComment, error) {
	var comments []DocComment
	err := s.client.GetRequestV2(s.endpoints.Default(channelID, docID), &comments)
	if err != nil {
		return nil, errors.New("error getting doc comments: " + err.Error())
	}

	return comments, nil
}
