package guildedgo

import (
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	config := &Config{
		ServerID: "",
		Token:    "",
	}

	c := NewClient(config)

	c.on("ChatMessageCreated", func(client *Client, e any) {
		data, ok := e.(*ChatMessageCreated)
		if ok {
			fmt.Println("New message: ", data.Message.Content)
		}
	})

	c.Open()
}
