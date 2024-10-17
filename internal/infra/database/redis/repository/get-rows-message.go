package repository

import (
	"context"
)

func (r *MesssageRepository) GetMessageRows() (int64, error) {
	ctx := context.Background()

	totalMessages, err := r.rdb.LLen(ctx, "messages").Result()

	if err != nil {
		return 0, err
	}

	return totalMessages, nil

}
