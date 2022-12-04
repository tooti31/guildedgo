package endpoints

var (
	ChannelEndpoint = func() string {
		return GuildedApi + "/channels"
	}
	ChannelEndpointWithID = func(channelId string) string {
		return GuildedApi + "/channels/" + channelId
	}
	CreateMessageEndpoint = func(channelId string) string {
		return GuildedApi + "/channels/" + channelId + "/messages"
	}
	GetChannelMessagesEndpoint = func(channelId string) string {
		return GuildedApi + "/channels/" + channelId + "/messages"
	}
	GetChannelMessageEndpoint = func(channelId string, messageId string) string {
		return GuildedApi + "/channels/" + channelId + "/messages/" + messageId
	}
	UpdateChannelMessageEndpoint = func(channelId string, messageId string) string {
		return GuildedApi + "/channels/" + channelId + "/messages/" + messageId
	}
)
