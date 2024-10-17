package repository

import (
	"context"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetUsersRows() (int64, error) {
	ctx := context.Background()

	totalUsers, err := r.rdb.LLen(ctx, "users").Result()

	if err != nil {
		return 0, err
	}

	divi := totalUsers / entity.PerPage

	return divi, nil

}
