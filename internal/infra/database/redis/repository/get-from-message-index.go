package repository

import (
	"context"
	"encoding/json"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetFromMessageIndex() (*[]entity.Message, error) {
	ctx := context.Background()
	startM = startM - entity.PerPage
	stopM = stopM - entity.PerPage
	multi := startM % entity.PerPage

	if multi == 0 {
		startM++
	}

	if startM < 0 {
		startM = 0
	}

	if stopM < 0 {
		return &[]entity.Message{}, nil
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

	if len(payloads) == 0 {
		return &[]entity.Message{}, nil
	}

	return &payloads, nil
}
