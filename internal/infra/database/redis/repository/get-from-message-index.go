package repository

import (
	"context"
	"encoding/json"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetFromMessageIndex() (*[]entity.Message, error) {
	ctx := context.Background()
	var start int64
	var stop int64
	totalMessages, err := r.rdb.LLen(ctx, "messages").Result()
	if err != nil {
		return nil, err
	}

	multi := totalMessages % entity.PerPage
	start = (totalMessages - (entity.StartMIndex)*entity.PerPage) - 1

	if multi == 0 {
		stop = (start + entity.PerPage) - 1
	} else {
		stop = (start + entity.PerPage) - 2
	}

	if start < 0 {
		start = 0
		stop++
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
