package server

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase/dto"
)

var (
	broadcast     = make(chan dto.Payload)
	users         = make(map[string]User)
	messageExists = make(map[*websocket.Conn]bool)
	mu            sync.Mutex

	ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
)
