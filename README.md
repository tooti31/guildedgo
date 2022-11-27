# guildedgo

A guilded.gg library in Go

## Getting started

```cmd
go get github.com/itschip/guildedgo
```

### Example

```go
func main() {
    serverID := GetEnv("SERVER_ID")
	token := GetEnv("TOKEN")

	config := &Config{
		ServerID: serverID,
		Token:    token,
	}

	c := NewClient(config)

	c.on("ChatMessageCreated", func(client *Client, e any) {
		data, ok := e.(*ChatMessageCreated)
		if ok {
			if data.Message.Content == "!ping" {
				client.Channel.SendMessage(data.Message.ChannelID, &MessageObject{
					Content: "Pong!",
				})
			}
		}
	})

	c.Open()
}
```
