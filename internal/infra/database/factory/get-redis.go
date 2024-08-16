package factory

import (
	"log"
	"strconv"

	connectionRedis "github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/redis/connection"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func (f *Factory) GetConRedis() (*redis.Client, error) {
	viper.AutomaticEnv()

	hostRedisDocker := viper.GetString("HOST_REDIS_DOCKER")
	hostRedis := f.Conf.HostRedis

	if hostRedisDocker != "" {
		hostRedis = hostRedisDocker
	}

	portR, errs := strconv.Atoi(f.Conf.PortRedis)
	if errs != nil {
		log.Fatalf("Invalid port: %v", errs)
	}

	return connectionRedis.ConnectingRedis(hostRedis, portR, f.Conf.PassRedis), nil
}
