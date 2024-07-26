package repository

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) SaveMessage(msg *entity.Message) error {
	ctx := context.Background()
	data, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	if strings.TrimSpace(msg.Message) != "" {
		err = r.rdb.RPush(ctx, "messages", data).Err()
		if err != nil {
			return err
		}
	}

	return nil
}
