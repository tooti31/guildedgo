# guildedgo

A guilded.gg library in Go

## Getting started

```cmd
go get github.com/itschip/guildedgo
```

### Example

```go
package main

import (
        "fmt"

        "github.com/itschip/guildedgo"
)

func main() {
        guildedClient := guildedgo.NewClient(&guildedgo.Config{
                Token:    "YOUR_TOKEN",
                ServerID: "YOUR_SERVER_ID",
        })

        // Listen to the ChatMessageCreated event
        guildedClient.On("ChatMessageCreated", func(client *guildedgo.Client, v any) {
                data, ok := v.(*guildedgo.ChatMessageCreated)

                if ok {
                        fmt.Println(data.Message.Content)

                        if data.Message.Content == "!ping" {
                                guildedClient.Channel.SendMessage(data.Message.ChannelID, &guildedgo.MessageObject{
                                        Content: "pong!",
                                })
                        }

                }
        })

        // Open socket
        guildedClient.Open()
}
```

## TODO

- [x] Channels
- [x] Servers
- [x] Messaging
- [x] Members
- [x] Member bans
- [x] Forums
- [ ] Forum comments
- [ ] List items
- [ ] Docs
- [ ] Calendar events
- [ ] Reactions
- [ ] Server XP
- [ ] Social links
- [ ] Group membership
- [ ] Role membership
- [ ] Webhooks
- [ ] Emote

```

```
