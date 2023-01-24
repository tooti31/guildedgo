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

	e, err := c.Calendar.GetEvents("", &GetEventsOptions{
		Limit: 1,
	})
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(e)

	c.Open()
}
