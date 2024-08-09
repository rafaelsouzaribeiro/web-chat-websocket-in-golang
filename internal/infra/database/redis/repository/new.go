package repository

import (
	"github.com/redis/go-redis/v9"
)

type MesssageRepository struct {
	rdb *redis.Client
}

var (
	startM int64 = 20
	stopM  int64 = 1
	startU int64 = 20
	stopU  int64 = 1
)

func NewMessageRedisRepository(db *redis.Client) *MesssageRepository {
	return &MesssageRepository{
		rdb: db,
	}
}
