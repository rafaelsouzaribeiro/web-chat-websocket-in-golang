package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

var inter int64 = 0

func (r *MesssageRepository) GetFromUsersIndex() (*[]entity.Message, error) {
	ctx := context.Background()

	startU = startU + entity.PerPage + 1
	stopU = stopU + entity.PerPage

	messages, err := r.rdb.LRange(ctx, "users", startU, stopU).Result()
	if err != nil {
		return nil, err
	}

	var payloads []entity.Message
	for _, msg := range messages {
		var payload entity.Message
		if err := json.Unmarshal([]byte(msg), &payload); err == nil {
			payload.Username = fmt.Sprintf("%s,%d", payload.Username, inter)
			payloads = append(payloads, payload)
		}
		inter++
	}

	return &payloads, nil
}
