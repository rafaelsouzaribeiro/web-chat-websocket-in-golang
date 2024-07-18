package repository

import (
	"context"
	"encoding/json"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) SaveMessage(msg *entity.Message) error {
	ctx := context.Background()
	data, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	err = r.rdb.RPush(ctx, "messages", data).Err()
	if err != nil {
		return err
	}

	return nil
}
