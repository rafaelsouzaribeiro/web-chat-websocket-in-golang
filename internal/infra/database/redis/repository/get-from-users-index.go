package repository

import (
	"context"
	"encoding/json"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetFromUsersIndex() (*[]entity.Message, error) {
	ctx := context.Background()
	EndIndex := entity.StartUIndex - 1
	entity.StartUIndex = EndIndex - 19
	if entity.StartUIndex < 0 {
		entity.StartUIndex = 0
	}

	messages, err := r.rdb.LRange(ctx, "users", entity.StartUIndex, EndIndex).Result()
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
