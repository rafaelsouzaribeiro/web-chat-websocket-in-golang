package client

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase/dto"
)

type Client struct {
	host    string
	port    int
	pattern string
	Conn    *websocket.Conn
	Channel chan dto.Payload
}

func NewClient(host, pattern string, port int) *Client {
	return &Client{
		host:    host,
		port:    port,
		pattern: pattern,
	}
}

func (client *Client) Connect() {

	url := fmt.Sprintf("ws://%s:%d/%s", client.host, client.port, client.pattern)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	client.Conn = conn
}

func (client *Client) Send(username, message string, currentTime time.Time, types string) {
	errs := client.Conn.WriteJSON(dto.Payload{Username: username, Message: message, Type: types, Time: currentTime})
	if errs != nil {
		log.Println("Error writing message:", errs)
		return
	}
}

func (client *Client) Listen() {
	for {
		var msg dto.Payload
		err := client.Conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		client.Channel <- msg
	}
}

func (client *Client) SendOne(username, message string, currentTime time.Time, types string, channel chan<- dto.Payload) {
	errs := client.Conn.WriteJSON(dto.Payload{Username: username, Message: message, Type: types, Time: currentTime})
	if errs != nil {
		log.Println("Error writing message:", errs)
		return
	}

	for {
		var msg dto.Payload
		err := client.Conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		channel <- msg
	}
}
