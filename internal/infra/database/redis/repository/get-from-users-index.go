package repository

import (
	"context"
	"encoding/json"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetFromUsersIndex() (*[]entity.Message, error) {
	ctx := context.Background()

	startM := float64(entity.PerPage) * (entity.StartUIndex - 2)
	stopM := startM + float64(entity.PerPage)

	var div int64
	if int64(entity.StartUIndex-2) != 0 {
		div = entity.PerPage % (int64(entity.StartUIndex) - 2)
	}

	if div == 0 {
		stopM = stopM - 1
	}

	if int64(startM) < 0 {
		startM = 0
	}

	if int64(stopM) <= 0 {
		return &[]entity.Message{}, nil
	}

	println(int64(startM), int64(stopM), (int64(entity.StartUIndex) - 2), div)
	messages, err := r.rdb.LRange(ctx, "users", int64(startM), int64(stopM)).Result()
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
