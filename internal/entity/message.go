package entity

import (
	"time"
)

var (
	StartMIndex = int64(-20)
	StartUIndex = int64(-20)
	PerPage     = int64(20)
)

type Message struct {
	Message  string
	Username string
	Type     string
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
