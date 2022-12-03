package endpoints

import (
	"github.com/itschip/guildedgo/internal"
)

var (
	CreateChannelEndpoint = func() string {
		return internal.GuildedApi + "/channels"
	}
	GetChannelEndpoint = func(channelId string) string {
		return internal.GuildedApi + "/channels/" + channelId
	}
	CreateMessageEndpoint = func(channelId string) string {
		return internal.GuildedApi + "/channels/" + channelId + "/messages"
	}
	GetChannelMessagesEndpoint = func(channelId string) string {
		return internal.GuildedApi + "/channels/" + channelId + "/messages"
	}
	GetChannelMessageEndpoint = func(channelId string, messageId string) string {
		return internal.GuildedApi + "/channels/" + channelId + "/messages/" + messageId
	}
	UpdateChannelMessageEndpoint = func(channelId string, messageId string) string {
		return internal.GuildedApi + "/channels/" + channelId + "/messages/" + messageId
	}
)
