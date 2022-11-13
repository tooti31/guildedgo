package guildedgo

func (c *Client) on(event string, callback func(client *Client, v *interface{})) {
	c.Events[event] = Event{
		Callback: callback,
		Type:     v,
	}
}
