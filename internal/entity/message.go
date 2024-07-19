package entity

import (
	"time"
)

var (
	StartMIndex = int64(-20)
	StartUIndex = int64(-20)
)

type Message struct {
	Message  string    `json:"message"`
	Username string    `json:"username"`
	Type     string    `json:"type"`
	Time     time.Time `json:"time"`
}

func NewMessage(message, username, t string, time time.Time) *Message {
	return &Message{
		Message:  message,
		Username: username,
		Type:     t,
		Time:     time,
	}
}
