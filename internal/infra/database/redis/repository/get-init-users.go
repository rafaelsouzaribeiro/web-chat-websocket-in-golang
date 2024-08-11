package repository

import (
	"context"
	"encoding/json"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetInitUsers() (*[]entity.Message, error) {
	ctx := context.Background()

	totalMessages, err := r.rdb.LLen(ctx, "users").Result()
	if err != nil {
		return nil, err
	}

	multi := totalMessages % entity.PerPage
	stopU = totalMessages

	if multi == 0 && totalMessages > ((entity.PerPage)-1) {
		startU = (totalMessages - entity.PerPage) - 1
	} else {
		startU = (totalMessages - entity.PerPage) - 2
	}

	if startU < 0 {
		startU = 0
	}

	messages, err := r.rdb.LRange(ctx, "users", startU, stopU).Result()
	if err != nil {
		return nil, err
	}

	var payloads []entity.Message
	for _, msg := range messages {
		var payload entity.Message
		if err := json.Unmarshal([]byte(msg), &payload); err == nil {
			payloads = append(payloads, payload)
		}

	}

	startU = 1
	stopU = entity.PerPage

	return &payloads, nil

}
