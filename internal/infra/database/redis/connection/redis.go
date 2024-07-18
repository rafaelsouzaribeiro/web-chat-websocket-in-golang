package connection

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func ConnectingRedis(host string, port int, password string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
	})

	return rdb
}
