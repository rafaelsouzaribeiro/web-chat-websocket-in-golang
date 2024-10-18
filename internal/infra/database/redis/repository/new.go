package repository

import (
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/factory"
	"github.com/redis/go-redis/v9"
)

type MesssageRepository struct {
	rdb *redis.Client
}

func NewMessageRedisRepository(db *factory.Iconnection) *MesssageRepository {
	return &MesssageRepository{
		rdb: db.Redis,
	}
}
