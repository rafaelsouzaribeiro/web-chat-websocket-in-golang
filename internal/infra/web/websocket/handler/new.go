package handler

import (
	"github.com/gorilla/websocket"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase"
)

type User struct {
	conn     *websocket.Conn
	username string
	id       string
}

type MessageHandler struct {
	messageUseCase *usecase.MessageUsecase
}

func NewMessageHandler(messageUseCase *usecase.MessageUsecase) *MessageHandler {
	return &MessageHandler{
		messageUseCase: messageUseCase,
	}
}
