package repository

import (
	"context"
	"encoding/json"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetFromMessageIndex() (*[]entity.Message, error) {
	ctx := context.Background()
	EndIndex := entity.StartMIndex - 1
	entity.StartMIndex = EndIndex - 19
	if entity.StartMIndex < 0 {
		entity.StartMIndex = 0
	}

	messages, err := r.rdb.LRange(ctx, "messages", entity.StartMIndex, EndIndex).Result()
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
