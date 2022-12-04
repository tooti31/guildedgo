package endpoints

var (
	MemberEndpoint = func(serverId string) string {
		return GuildedApi + "/servers/" + serverId + "/members"
	}
	ServerMemberEndpoint = func(serverId string, userId string) string {
		return GuildedApi + "/servers/" + serverId + "/members/" + userId
	}
	MemberNicknameEndpoint = func(serverId string, userId string) string {
		return ServerMemberEndpoint(serverId, userId) + "/nickname"
	}
	MemberBanEndpoint = func(serverId string, userId string) string {
		return GuildedApi + "/servers/" + serverId + "/bans/" + userId
	}
)
