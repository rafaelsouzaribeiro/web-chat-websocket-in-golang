package entity

import (
	"time"
)

var (
	StartMIndex = int64(-20)
	StartUIndex = int64(-20)
	PerPage     = int64(20)
	KeySpace    = "chat"
	PageM       = int64(1)
	TotalM      = int64(1)
	PointerM    = int64(1)
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
