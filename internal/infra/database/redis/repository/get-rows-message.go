package repository

import (
	"context"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetMessageRows() (float64, error) {
	ctx := context.Background()

	totalMessages, err := r.rdb.LLen(ctx, "messages").Result()

	if err != nil {
		return 0, err
	}

	divi := float64(totalMessages / entity.PerPage)

	return divi, nil

}
