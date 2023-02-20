package guildedgo

import (
	"fmt"
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

	commands := &CommandsBuilder{
		Commands: []Command{
			{
				CommandName: "!test",
				Action: func(client *Client, v *ChatMessageCreated) {
					client.Channel.SendMessage(v.Message.ChannelID, &MessageObject{
						Content: "Test",
					})

					fmt.Println("Test working")
				},
			},
			{
				CommandName: "!party",
				Action: func(client *Client, v *ChatMessageCreated) {
					client.Channel.SendMessage(v.Message.ChannelID, &MessageObject{
						Content: "Yeah!!! Let's party",
					})

					fmt.Println("Party working")
				},
			},
		},
	}

	c.CommandService.AddCommands(commands)

	c.Open()
}
