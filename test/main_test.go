package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase/dto"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/test/client"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/test/server"
	"github.com/stretchr/testify/assert"
)

func BenchmarkUser(b *testing.B) {
	server.Once.Do(server.StartServer)

	channel := make(chan dto.Payload, b.N)
	var messages []dto.Payload
	client := client.NewClient("0.0.0.0", "ws", 8080)
	client.Channel = channel

	for i := 0; i < b.N; i++ {
		server.Count++
		go func(count int) {
			client.Connect()
			go client.Listen()
			client.Send(fmt.Sprintf("Client %d", server.Count), "", time.Now(), "")
			client.Send(fmt.Sprintf("Client %d", server.Count), fmt.Sprintf("Message %d", server.Count), time.Now(), "message")
		}(server.Count)
	}

	timeout := time.After(5 * time.Second)
loop:
	for {
		select {
		case msg := <-channel:
			messages = append(messages, msg)
		case <-timeout:
			break loop
		}
	}

	for i, msg := range messages {
		if msg.Type == "message" {
			user := fmt.Sprintf("<strong>Client %d</strong> ", i)

			if user == msg.Username {
				assert.Contains(b, msg.Username, fmt.Sprintf("Client %d", i))
			}
		} else {
			assert.Contains(b, msg.Username, "info")
			connected := fmt.Sprintf("User <strong>Client %d</strong> connected", i)

			if connected == msg.Message {
				assert.Contains(b, msg.Message, fmt.Sprintf("Client %d", i))
			}
		}
		b.Logf("%d %s %s %s", i, msg.Message, msg.Username, msg.Type)
	}
}
