package repository

import (
	"github.com/redis/go-redis/v9"
)

type MesssageRepository struct {
	rdb *redis.Client
}

var (
	startM int64 = 1
	stopM  int64 = 20
	startU int64 = 1
	stopU  int64 = 20
)

func NewMessageRedisRepository(db *redis.Client) *MesssageRepository {
	return &MesssageRepository{
		rdb: db,
	}
}
