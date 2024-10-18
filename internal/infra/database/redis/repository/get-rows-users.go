package repository

import (
	"context"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetUsersRows() (float64, error) {
	ctx := context.Background()

	totalUsers, err := r.rdb.LLen(ctx, "users").Result()

	if err != nil {
		return 0, err
	}

	divi := float64(totalUsers / entity.PerPage)

	return divi, nil

}
