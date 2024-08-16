package di

import (
	cassandra "github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/cassandra/repository"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/factory"
	redis "github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/redis/repository"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase"
)

func NewUseCase(db *factory.Iconnection) *usecase.MessageUsecase {

	if db.Gocql != nil {
		repository := cassandra.NewMessageCassandraRepository(db)
		return usecase.NewMessageUseCase(repository)
	} else if db.Redis != nil {
		repository := redis.NewMessageRedisRepository(db)
		return usecase.NewMessageUseCase(repository)
	}

	return nil
}
