package guildedgo

import (
	"errors"
	"strconv"
)

type reactionEndpoints struct{}

func (e *reactionEndpoints) Default(channelId string, contentId string, emoteId int) string {
	return guildedApi + "/channels/" + channelId + "/content/" + contentId + "/emotes/" + strconv.Itoa(emoteId)
}

func (e *reactionEndpoints) Topic(channelId string, topicId int, emoteId int) string {
	return guildedApi + "/channels/" + channelId + "/topics/" + strconv.Itoa(topicId) + "/emotes/" + strconv.Itoa(emoteId)
}

type ReactionService interface {
	AddReactionEmote(channelId string, contentId string, emoteId int) error
	DeleteReactionEmote(channelId string, contentId string, emoteId int) error
	AddTopicReactionEmote(channelId string, topicId int, emoteId int) error
	DeleteTopicReactionEmote(channelId string, topicId int, emoteId int) error
}

type reactionService struct {
	client    *Client
	endpoints *reactionEndpoints
}

var _ ReactionService = &reactionService{}

func (service *reactionService) AddReactionEmote(channelId string, contentId string, emoteId int) error {
	_, err := service.client.PutRequest(service.endpoints.Default(channelId, contentId, emoteId), nil)
	if err != nil {
		return errors.New("Failed to add reaction emote: " + err.Error())
	}

	return nil
}

func (service *reactionService) DeleteReactionEmote(channelId string, contentId string, emoteId int) error {
	_, err := service.client.DeleteRequest(service.endpoints.Default(channelId, contentId, emoteId))
	if err != nil {
		return errors.New("Failed to delete reaction emote: " + err.Error())
	}

	return nil
}

func (service *reactionService) AddTopicReactionEmote(channelId string, topicId int, emoteId int) error {
	_, err := service.client.PutRequest(service.endpoints.Topic(channelId, topicId, emoteId), nil)
	if err != nil {
		return errors.New("Failed to add topic reaction emote: " + err.Error())
	}

	return nil
}

func (service *reactionService) DeleteTopicReactionEmote(channelId string, topicId int, emoteId int) error {
	_, err := service.client.DeleteRequest(service.endpoints.Topic(channelId, topicId, emoteId))
	if err != nil {
		return errors.New("Failed to add topic reaction emote: " + err.Error())
	}

	return nil
}
