package guildedgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func (c *Client) Open() {
	header := http.Header{}
	header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	conn, _, err := websocket.DefaultDialer.Dial("wss://www.guilded.gg/websocket/v1", header)
	if err != nil {
		log.Fatalln("Failed to connect to websocket: ", err.Error())
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}()

	_, m, err := conn.ReadMessage()
	if err != nil {
		log.Fatalln("Failed to read message: ", err.Error())
	}

	m = bytes.TrimSpace(bytes.Replace(m, newline, space, -1))
	//	fmt.Println(string(m))

	listening := make(chan struct{})

	go func() {
		defer close(listening)

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err.Error())
				return
			}

			c.onEvent(msg)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-listening:
			return

		case t := <-ticker.C:
			err := conn.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write", err)
				return
			}
		case <-interrupt:
			log.Println("Interrupt")

			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close", err)
				return
			}

			select {
			case <-listening:
			case <-time.After(time.Second):
			}

			return
		}
	}
}

var interfaces = make(map[string]interface{})

func init() {
	interfaces["ChatMessageCreated"] = &ChatMessage{}
}

type SocketMessage[T any] struct {
	T string `json:"t"`
	S string `json:"s"`
	D string `json:"d"`
}

func (c *Client) onEvent(msg []byte) {
	fmt.Println(string(msg))

	var message SocketMessage[string]

	err := json.Unmarshal(msg, &message)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("unmarshal message")
	fmt.Println(message)

	fmt.Println("message data")
	fmt.Println([]byte(message.D))

	eventInterface := interfaces[message.T]

	err = json.Unmarshal([]byte(message.D), &eventInterface)
	if err != nil {
		log.Println("Failed to umarshal chat message event")
	}

	fmt.Println("eventInterface moment")
	fmt.Println(eventInterface)

	/*c.Events[message.T](c, &eventInterface) */
}
