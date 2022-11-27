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

	c.on("ChatMessageCreated", func(client *Client, e any) {
		data, ok := e.(*ChatMessageCreated)
		if ok {
			if data.Message.Content == "!ping" {
				client.Channel.SendMessage(data.Message.ChannelID, &MessageObject{
					Content: "Pong!",
				})
			}
		}
	})

	c.Open()
}
