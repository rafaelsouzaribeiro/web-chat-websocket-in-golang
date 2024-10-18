package repository

import (
	"context"
	"encoding/json"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetFromUsersIndex() (*[]entity.Message, error) {
	ctx := context.Background()

	stopM := entity.PerPage * (entity.StartUIndex - 1)
	startM := stopM - entity.PerPage

	if stopM <= 0 {
		return &[]entity.Message{}, nil
	}

	if startM < 0 {
		startM = 0
	}
	println(startM, stopM, entity.StartUIndex)
	messages, err := r.rdb.LRange(ctx, "users", startM, stopM).Result()
	if err != nil {
		return nil, err
	}

	payloads := make([]entity.Message, len(messages))
	var inter int = len(messages) - 1
	for _, msg := range messages {
		var payload entity.Message
		if err := json.Unmarshal([]byte(msg), &payload); err == nil {
			payloads[inter] = payload
		}
		inter--
	}

	if len(payloads) == 0 {
		return &[]entity.Message{}, nil
	}

	return &payloads, nil
}
