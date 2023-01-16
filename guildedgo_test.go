package guildedgo

import (
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

	err := c.Forums.LockForumTopic("", 0)
	if err != nil {
		log.Println(err)
	}

	c.Open()
}
