package endpoints

var ForumTopicEndpoint = func(channelId string) string {
	return GuildedApi + "/channels/" + channelId + "/topics"
}
