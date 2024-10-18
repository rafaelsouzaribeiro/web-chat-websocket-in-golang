package repository

import (
	"context"
	"encoding/json"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetFromMessageIndex() (*[]entity.Message, error) {
	ctx := context.Background()
	stopM := float64(entity.PerPage) * entity.StartMIndex
	startM := stopM - float64(entity.PerPage)

	if startM < 0 {
		startM = 0
	}

	messages, err := r.rdb.LRange(ctx, "messages", int64(startM), int64(stopM)).Result()
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
