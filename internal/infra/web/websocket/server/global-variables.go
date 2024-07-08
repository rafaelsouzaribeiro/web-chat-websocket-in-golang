package server

import (
	"context"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase/dto"
	"github.com/redis/go-redis/v9"
)

var (
	broadcast     = make(chan dto.Payload)
	users         = make(map[string]User)
	messageExists = make(map[*websocket.Conn]bool)
	mu            sync.Mutex

	ctx = context.Background()
	rdb *redis.Client
)

const (
	perPage = 20
)
