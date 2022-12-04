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

	c.Command("!updatenick", func(client *Client, v *ChatMessageCreated) {
		fmt.Println("Created by: ", v.Message.CreatedBy)
		nick, err := client.Members.UpdateMemberNickname(v.Message.CreatedBy, "coolnick")
		if err != nil {
			log.Println(err.Error())
		}

		fmt.Println("new nick: ", nick)
	})

	c.Open()
}
