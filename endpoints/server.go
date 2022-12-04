package endpoints

const (
	GuildedApi = "https://www.guilded.gg/api/v1"
)

var ServerEndpoint = func(serverId string) string {
	return GuildedApi + "/servers/" + serverId
}
