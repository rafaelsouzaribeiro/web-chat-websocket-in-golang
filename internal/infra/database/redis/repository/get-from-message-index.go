package repository

import (
	"context"
	"encoding/json"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetFromMessageIndex() (*[]entity.Message, error) {
	ctx := context.Background()

	startM := float64(entity.PerPage) * (entity.StartMIndex - 2)
	stopM := startM + float64(entity.PerPage)

	if int64(entity.StartMIndex) == 1 {
		stopM = stopM - 1
	}

	if int64(startM) < 0 {
		startM = 0
	}

	if int64(stopM) <= 0 {
		return &[]entity.Message{}, nil
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
