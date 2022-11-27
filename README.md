# guildedgo

A guilded.gg library in Go

## Getting started

```cmd
go get github.com/itschip/guildedgo
```

### Example

```go
func main() {
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
```
