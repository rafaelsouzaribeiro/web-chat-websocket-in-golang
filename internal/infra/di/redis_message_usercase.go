package di

import (
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/redis/repository"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase"
	"github.com/redis/go-redis/v9"
)

func NewMessageRedisUseCase(db *redis.Client) *usecase.MessageUsecase {
	repository := repository.NewMessageRedisRepository(db)
	return usecase.NewMessageUseCase(repository)
}
