package repository

import (
	"github.com/redis/go-redis/v9"
)

type MesssageRepository struct {
	rdb *redis.Client
}

func NewMessageRedisRepository(db *redis.Client) *MesssageRepository {
	return &MesssageRepository{
		rdb: db,
	}
}
