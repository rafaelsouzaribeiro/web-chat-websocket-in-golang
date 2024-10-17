package repository

import (
	"context"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetMessageRows() (int64, error) {
	ctx := context.Background()

	totalMessages, err := r.rdb.LLen(ctx, "messages").Result()

	if err != nil {
		return 0, err
	}

	divi := totalMessages / entity.PerPage

	return divi, nil

}
