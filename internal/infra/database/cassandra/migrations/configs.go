package migrations

import (
	"strings"

	"github.com/gocql/gocql"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/configs"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/cassandra/connection"
	"github.com/spf13/viper"
)

func SetVariables() (*gocql.Session, error) {

	viper.AutomaticEnv()
	hostsDocker := strings.Split(viper.GetString("HOST_CASSANDRA_DOCKER"), ",")

	Conf, err := configs.LoadConfig("./")

	if err != nil {
		panic(err)
	}

	user := Conf.UserCassaandra
	password := Conf.PassCassaandra
	hosts := strings.Split(Conf.HostCassaandra, ",")

	if len(hostsDocker) > 0 {
		hosts = hostsDocker
	}

	con, err := connection.NewCassandraConnection(hosts, user, password)

	if err != nil {
		return nil, err
	}

	return con, nil

}
