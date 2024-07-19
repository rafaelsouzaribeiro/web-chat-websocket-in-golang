package repository

import (
	"github.com/redis/go-redis/v9"
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
