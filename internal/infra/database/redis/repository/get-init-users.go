package repository

import (
	"context"
	"encoding/json"
	"sort"
	"time"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetInitUsers() (*[]entity.Message, error) {
	ctx := context.Background()

	totalMessages, err := r.rdb.LLen(ctx, "users").Result()
	if err != nil {
		return nil, err
	}
	stopM := totalMessages
	startM := stopM - entity.PerPage + 1

	if startM <= 0 {
		startM = 0
	}

	messages, err := r.rdb.LRange(ctx, "users", int64(startM), int64(stopM)).Result()
	if err != nil {
		return nil, err
	}

	var payloads []entity.Message
	for _, msg := range messages {
		var payload entity.Message
		if err := json.Unmarshal([]byte(msg), &payload); err == nil {
			payload.Time = time.Now()
			payloads = append(payloads, payload)
		}

	}

	sort.Slice(payloads, func(i, j int) bool {
		return payloads[i].Time.After(payloads[j].Time)
	})

	return &payloads, nil

}
