package endpoints

var GroupMemberEndpoint = func(groupId string, userId string) string {
	return GuildedApi + "/groups/" + groupId + "/members/" + userId
}
