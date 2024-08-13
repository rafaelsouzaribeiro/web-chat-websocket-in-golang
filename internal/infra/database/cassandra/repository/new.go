package repository

import (
	"sync"

	"github.com/gocql/gocql"
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

func NewMessageCassandraRepository(db *gocql.Session) *MesssageRepository {
	return &MesssageRepository{
		cql: db,
	}
}
