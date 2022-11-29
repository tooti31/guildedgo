package guildedgo

func (c *Client) On(event string, callback func(client *Client, v any)) {
	c.Events[event] = Event{
		Callback: callback,
	}
}

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
