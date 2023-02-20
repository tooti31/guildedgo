package guildedgo

import "fmt"

// On listens to any event
func (c *Client) On(event string, callback func(client *Client, v any)) {
	fmt.Println("On called", event)
	c.Events[event] = Event{
		Callback: callback,
	}
}

// Command listens to ChatMessageCreated and fires a func when the message content matches the command
func (c *Client) Command(cmd string, callback func(client *Client, v *ChatMessageCreated)) {
	fmt.Println("Command called", cmd)
	c.On("ChatMessageCreated", func(client *Client, v any) {
		data, ok := v.(*ChatMessageCreated)
		if ok {
			if data.Message.Content == cmd {
				callback(client, data)
			}
		}
	})
}
