package entity

import (
	"time"
)

var (
	StartMIndex float64
	StartUIndex float64
	PerPage     = int64(20)
	KeySpace    = "chat"
)

type Message struct {
	Message  string
	Username string
	Type     string
	Pages    int
	Time     time.Time
}

func NewMessage(message, username, t string, time time.Time) *Message {
	return &Message{
		Message:  message,
		Username: username,
		Type:     t,
		Time:     time,
	}
}
