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

	if totalMessages > entity.PerPage {
		entity.StartUIndex = totalMessages - entity.PerPage
	}

	messages, err := r.rdb.LRange(ctx, "users", entity.StartUIndex, -1).Result()
	if err != nil {
		return nil, err
	}

	var payloads []entity.Message
	for i, msg := range messages {
		if i == 0 {
			continue
		}
		var payload entity.Message
		if err := json.Unmarshal([]byte(msg), &payload); err == nil {
			payloads = append(payloads, payload)
		}

	}

	return &payloads, nil

}
