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
	if multi == 0 && totalMessages > 19 {
		startU = startU + 20
		stopU = stopU + 10
	}

	messages, err := r.rdb.LRange(ctx, "users", (startU * -1), (stopU * -1)).Result()
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

	return &payloads, nil

}
