package dto

import "time"

type Payload struct {
	Message  string    `json:"message"`
	Username string    `json:"username"`
	Type     string    `json:"type"`
	Time     time.Time `json:"time"`
}
