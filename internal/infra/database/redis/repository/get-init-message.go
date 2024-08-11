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

	multi := totalMessages % entity.PerPage
	stopM = totalMessages

	if multi == 0 && totalMessages > 19 {
		startM = (totalMessages - entity.PerPage) - 1
	} else {
		startM = (totalMessages - entity.PerPage) - 2
	}

	if startM < 0 {
		startM = 0
	}

	messages, err := r.rdb.LRange(ctx, "messages", startM, stopM).Result()
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

	startM = 1
	stopM = entity.PerPage

	return &payloads, nil

}
