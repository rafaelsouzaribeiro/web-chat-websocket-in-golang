package repository

import (
	"context"
	"encoding/json"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetFromUsersIndex() (*[]entity.Message, error) {
	ctx := context.Background()
	var start int64
	var stop int64
	totalMessages, err := r.rdb.LLen(ctx, "users").Result()
	if err != nil {
		return nil, err
	}

	multi := totalMessages % 20
	start = (totalMessages - (entity.StartUIndex)*20) - 1

	if multi == 0 {
		stop = (start + 20) - 1
	} else {
		stop = (start + 20) - 2
	}

	if start < 0 {
		start = 0
	}

	if stop < 0 {
		return &[]entity.Message{}, nil
	}

	if stop < 0 {
		return &[]entity.Message{}, nil
	}

	messages, err := r.rdb.LRange(ctx, "users", start, stop).Result()
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
