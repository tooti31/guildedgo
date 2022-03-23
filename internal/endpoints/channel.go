package endpoints

import (
	"github.com/itschip/guildedgo/internal"
)

var (
	CreateMessageEndpoint = func(channelId string) string {
		return internal.GuildedApi + "/channels/" + channelId + "/messages"
	}
)
