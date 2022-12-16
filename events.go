package guildedgo

// On listens to any event
func (c *Client) On(event string, callback func(client *Client, v any)) {
	c.Events[event] = Event{
		Callback: callback,
	}
}

// Command listens to ChatMessageCreated and fires a func when the message content matches the command
func (c *Client) Command(cmd string, callback func(client *Client, v *ChatMessageCreated)) {
	c.On("ChatMessageCreated", func(client *Client, v any) {
		data, ok := v.(*ChatMessageCreated)
		if ok {
			if data.Message.Content == cmd {
				callback(client, data)
			}
		}
	})
}
