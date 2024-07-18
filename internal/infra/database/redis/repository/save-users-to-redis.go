package repository

import (
	"context"
	"encoding/json"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) SaveUsers(msg *entity.Message) error {
	ctx := context.Background()
	data, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	err = r.rdb.RPush(ctx, "users", data).Err()
	if err != nil {
		return err
	}

	return nil
}
