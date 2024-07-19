package repository

import (
	"github.com/redis/go-redis/v9"
)

var (
	StartMIndex = int64(-20)
	StartUIndex = int64(-20)
)

const PerPage = 20

type MesssageRepository struct {
	rdb *redis.Client
}

func NewMessageRepository(db *redis.Client) *MesssageRepository {
	return &MesssageRepository{
		rdb: db,
	}
}
