package factory

import (
	"strings"

	"github.com/gocql/gocql"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/cassandra/connection"
	"github.com/spf13/viper"
)

func (f *Factory) GetConCassandra() (*gocql.Session, error) {
	viper.AutomaticEnv()

	hostsDocker := strings.Split(viper.GetString("HOST_CASSANDRA_DOCKER"), ",")
	hosts := strings.Split(f.Conf.HostCassaandra, ",")

	if hostsDocker[0] != "" {
		hosts = hostsDocker
	}

	cassandra, err := connection.NewCassandraConnection(hosts, f.Conf.PassCassaandra, f.Conf.UserCassaandra)

	if err != nil {
		return nil, err
	}

	return cassandra, nil
}
