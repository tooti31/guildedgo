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

func (e *docCommentEndpoints) Get(channelID string, docID int, commentID int) string {
	return guildedApi + "/channels/" + channelID + "/docs/" + strconv.Itoa(docID) + "/comments/" + strconv.Itoa(commentID)
}

type DocCommentService interface {
	Create(channelID string, docID int, content string) (*DocComment, error)
	GetComments(channelID string, docID int) ([]DocComment, error)
	GetDocComment(channelID string, docID int, commentID int) (*DocComment, error)
	UpdateDocComment(channelID string, docID int, commentID int, content string) (*DocComment, error)
	DeleteDocComment(channelID string, docID int, commentID int) error
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

func (s *docCommentService) GetDocComment(channelID string, docID int, commentID int) (*DocComment, error) {
	var comment DocComment
	err := s.client.GetRequestV2(s.endpoints.Get(channelID, docID, commentID), &comment)
	if err != nil {
		return nil, errors.New("error getting doc comment: " + err.Error())
	}

	return &comment, nil
}

func (s *docCommentService) UpdateDocComment(channelID string, docID int, commentID int, content string) (*DocComment, error) {
	var comment DocComment
	err := s.client.PatchRequest(s.endpoints.Get(channelID, docID, commentID), content, &comment)
	if err != nil {
		return nil, errors.New("error updating doc comment: " + err.Error())
	}

	return &comment, nil
}

func (s *docCommentService) DeleteDocComment(channelID string, docID int, commentID int) error {
	_, err := s.client.DeleteRequest(s.endpoints.Get(channelID, docID, commentID))
	if err != nil {
		return errors.New("error deleting doc comment: " + err.Error())
	}

	return nil
}
