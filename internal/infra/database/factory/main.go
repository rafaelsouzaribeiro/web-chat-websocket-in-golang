package factory

import (
	"log"
	"strconv"
	"strings"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/configs"
	ConnectingCassandra "github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/cassandra/connection"
	ConnectingRedis "github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/redis/connection"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/di"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase"
	"github.com/spf13/viper"
)

const (
	Redis     = "redis"
	Cassandra = "cassandra"
)

func NewFactory(types string, Conf *configs.Conf) *usecase.MessageUsecase {
	viper.AutomaticEnv()

	switch types {
	case "redis":
		hostRedisDocker := viper.GetString("HOST_REDIS_DOCKER")
		hostRedis := Conf.HostRedis

		if hostRedisDocker != "" {
			hostRedis = hostRedisDocker
		}

		portR, errs := strconv.Atoi(Conf.PortRedis)
		if errs != nil {
			log.Fatalf("Invalid port: %v", errs)
		}

		redis := ConnectingRedis.ConnectingRedis(hostRedis, portR, Conf.PassRedis)
		di := di.NewMessageRedisUseCase(redis)

		return di
	case "cassandra":
		hostsDocker := strings.Split(viper.GetString("HOST_CASSANDRA_DOCKER"), ",")
		hosts := strings.Split(Conf.HostCassaandra, ",")

		if hostsDocker[0] != "" {
			hosts = hostsDocker
		}

		cassandra, err := ConnectingCassandra.NewCassandraConnection(hosts, Conf.PassCassaandra, Conf.UserCassaandra)

		if err != nil {
			panic(err)
		}

		di := di.NewMessageCassandraUseCase(cassandra)

		return di
	}

	return nil
}
