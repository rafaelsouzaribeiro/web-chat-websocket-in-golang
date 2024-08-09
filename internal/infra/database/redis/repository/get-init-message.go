package repository

import (
	"context"
	"encoding/json"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetInitMessages() (*[]entity.Message, error) {
	ctx := context.Background()

	totalMessages, err := r.rdb.LLen(ctx, "messages").Result()
	if err != nil {
		return nil, err
	}

	multi := totalMessages % 20
	if multi == 0 && totalMessages > 19 {
		startM = startM + 20
		stopM = stopM + 10
	}

	messages, err := r.rdb.LRange(ctx, "messages", (startM * -1), stopM*-1).Result()
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
