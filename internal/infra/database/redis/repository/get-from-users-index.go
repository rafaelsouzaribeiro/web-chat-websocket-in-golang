package repository

import (
	"context"
	"encoding/json"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetFromUsersIndex() (*[]entity.Message, error) {
	ctx := context.Background()

	startU = startU + entity.PerPage
	stopU = stopU + entity.PerPage

	if startU == 41 {
		startU++
	}

	messages, err := r.rdb.LRange(ctx, "users", startU, stopU).Result()
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

	return &payloads, nil
}
