package di

import (
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/redis/repository"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase"
	"github.com/redis/go-redis/v9"
)

func NewMessageUseCase(db *redis.Client) *usecase.MessageUsecase {
	repository := repository.NewMessageRepository(db)
	return usecase.NewMessageUseCase(repository)
}
