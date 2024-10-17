package repository

import (
	"context"
)

func (r *MesssageRepository) GetUsersRows() (int64, error) {
	ctx := context.Background()

	totalUsers, err := r.rdb.LLen(ctx, "users").Result()

	if err != nil {
		return 0, err
	}

	return totalUsers, nil

}
