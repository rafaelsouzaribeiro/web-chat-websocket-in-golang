package dto

import (
	"github.com/gorilla/websocket"
)

type Payload struct {
	Message  string `json:"message"`
	Username string `json:"username"`
	Id       string
	Conn     *websocket.Conn
	Type     string `json:"type"`
}
