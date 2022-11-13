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

	c.on("ChatMessageCreated", func(c *Client, cmc *interface{}) {
		fmt.Println("cmc", cmc)
	})

	c.Open()
}
