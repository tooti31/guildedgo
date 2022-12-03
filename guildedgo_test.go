package guildedgo

import (
	"fmt"
	"log"
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

	c.On("ChatMessageCreated", func(client *Client, e any) {
		data, ok := e.(*ChatMessageCreated)
		if ok {
			if data.Message.Content == "!ping" {
				client.Channel.SendMessage(data.Message.ChannelID, &MessageObject{
					Content: "Pong!",
				})
			}
		}
	})

	c.Command("!create", func(client *Client, v *ChatMessageCreated) {
		channel, err := c.Channel.GetChannel("a6bc58e4-9d63-4290-b674-d8f18ecb28fb")
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(channel)
	})

	c.Open()
}
