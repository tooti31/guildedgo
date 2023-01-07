package endpoints

import "fmt"

var ForumTopicEndpoint = func(channelId string) string {
	return GuildedApi + "/channels/" + channelId + "/topics"
}

var GetForumTopicEndpoint = func(channelId string, forumTopicId int) string {
	return GuildedApi + "/channels/" + channelId + "/topics/" + fmt.Sprint(forumTopicId)
}
