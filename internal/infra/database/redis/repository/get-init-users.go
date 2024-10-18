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

	stopM := totalMessages
	startM := stopM - entity.PerPage

	if stopM <= 0 {
		return &[]entity.Message{}, nil
	}

	if startM < 0 || startM < 20 {
		startM = 0
	}

	println(startM, stopM, entity.StartUIndex)
	messages, err := r.rdb.LRange(ctx, "users", startM, stopM).Result()
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
