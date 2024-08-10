package repository

import (
	"context"
	"encoding/json"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetFromMessageIndex() (*[]entity.Message, error) {
	ctx := context.Background()
	entity.StartMIndex = (entity.StartMIndex - 1) * entity.PerPage
	stop := entity.StartMIndex + entity.PerPage

	totalMessages, err := r.rdb.LLen(ctx, "messages").Result()
	if err != nil {
		return nil, err
	}

	if stop > totalMessages {
		stop = totalMessages
	}

	messages, err := r.rdb.LRange(ctx, "messages", (stop * -1), ((entity.StartMIndex * -1) - 1)).Result()
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
