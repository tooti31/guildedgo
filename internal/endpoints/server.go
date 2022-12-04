package endpoints

import "github.com/itschip/guildedgo/internal"

var ServerEndpoint = func(serverId string) string {
	return internal.GuildedApi + "/servers/" + serverId
}
