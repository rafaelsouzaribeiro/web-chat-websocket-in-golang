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

	multi := totalMessages % 20
	stopU = totalMessages

	if multi == 0 && totalMessages > 19 {
		startU = (totalMessages - 20) - 1
	} else {
		startU = (totalMessages - 20) - 2
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
	stopU = 20

	return &payloads, nil

}
