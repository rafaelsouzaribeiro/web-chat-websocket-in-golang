package handler

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase/dto"
)

var (
	connected     = make(chan dto.Payload)
	messages      = make(chan dto.Payload)
	users         = make(map[string]User)
	messageExists = make(map[*websocket.Conn]bool)
	mu            sync.Mutex

	StartMIndex = int64(-20)
	StartUIndex = int64(-20)
)

const (
	PerPage = 20
)
