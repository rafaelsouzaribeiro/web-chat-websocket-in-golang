package repository

import (
	"sync"

	"github.com/gocql/gocql"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/factory"
)

type MesssageRepository struct {
	cql *gocql.Session
}

type Pagination struct {
	Id    string
	Total int
	Page  int
}

var Once sync.Once

func NewMessageCassandraRepository(db *factory.Iconnection) *MesssageRepository {
	return &MesssageRepository{
		cql: db.Gocql,
	}
}
