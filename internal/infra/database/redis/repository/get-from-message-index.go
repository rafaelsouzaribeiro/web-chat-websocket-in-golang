package repository

import (
	"context"
	"encoding/json"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetFromMessageIndex() (*[]entity.Message, error) {
	ctx := context.Background()

	totalMessages, err := r.rdb.LLen(ctx, "messages").Result()
	if err != nil {
		return nil, err
	}

	start := (totalMessages - (entity.StartMIndex)*20) - 1

	stop := (start + 20) - 1

	if start < 0 {
		start = 0
	}

	if stop < 0 {
		return &[]entity.Message{}, nil
	}

	messages, err := r.rdb.LRange(ctx, "messages", start, stop).Result()
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
