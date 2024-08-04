package repository

import (
	"github.com/gocql/gocql"
)

type MesssageRepository struct {
	cql *gocql.Session
}

type Save struct {
	Total int
	Page  int
}

func NewMessageRepository(db *gocql.Session) *MesssageRepository {
	return &MesssageRepository{
		cql: db,
	}
}
