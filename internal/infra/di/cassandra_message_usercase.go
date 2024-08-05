package di

import (
	"github.com/gocql/gocql"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/cassandra/repository"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase"
)

func NewMessageCassandraUseCase(db *gocql.Session) *usecase.MessageUsecase {
	repository := repository.NewMessageCassandraRepository(db)
	return usecase.NewMessageUseCase(repository)
}
