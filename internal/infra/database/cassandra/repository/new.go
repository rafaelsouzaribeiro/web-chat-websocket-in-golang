package repository

import (
	"github.com/gocql/gocql"
)

type MesssageRepository struct {
	cql *gocql.Session
}

type Save struct {
	Id    gocql.UUID
	Total int
	Page  int
}

func NewMessageCassandraRepository(db *gocql.Session) *MesssageRepository {
	return &MesssageRepository{
		cql: db,
	}
}
