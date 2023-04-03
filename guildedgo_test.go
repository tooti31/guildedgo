package guildedgo

import (
	"testing"

	"github.com/itschip/guildedgo/internal"
)

func TestNewClient(t *testing.T) {
	serverID := internal.GetEnv("SERVER_ID")
	token := internal.GetEnv("TOKEN")

	config := &Config{
		ServerID: serverID,
		Token:    token,
	}

	c := NewClient(config)

	_, err := c.Channel.UpdateChannel("what", &UpdateChannelObject{
		Name: "whehe",
	})
	if err != nil {
		c.Channel.SendMessage("", &MessageObject{
			Content: "Failed to update channel",
		})
	}
}
