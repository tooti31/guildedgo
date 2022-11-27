package guildedgo

func (c *Client) on(event string, callback func(client *Client, v any)) {
	c.Events[event] = Event{
		Callback: callback,
	}
}
