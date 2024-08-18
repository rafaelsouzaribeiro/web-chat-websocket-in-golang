package factory

import (
	"github.com/gocql/gocql"
	"github.com/redis/go-redis/v9"
)

const (
	Redis     = "redis"
	Cassandra = "cassandra"
)

type Iconnection struct {
	Gocql *gocql.Session
	Redis *redis.Client
}
