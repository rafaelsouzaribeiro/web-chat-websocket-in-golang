package server

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func (server *Server) ConnectingRedis(host string, port int) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: "123mudar",
	})
}
