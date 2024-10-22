package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetInitUsers() (*[]entity.Message, error) {
	ctx := context.Background()

	totalUsers, err := r.rdb.LLen(ctx, "users").Result()
	if err != nil {
		return nil, err
	}
	stopM := totalUsers
	startM := stopM - entity.PerPage + 1

	if startM <= 0 {
		startM = 0
	}

	users, err := r.rdb.LRange(ctx, "users", startM, stopM).Result()
	if err != nil {
		return nil, err
	}

	var payloads []entity.Message
	for _, msg := range users {
		var payload entity.Message
		if err := json.Unmarshal([]byte(msg), &payload); err == nil {
			payload.Time = time.Now()
			payloads = append(payloads, payload)
		}

	}

	return &payloads, nil

}
