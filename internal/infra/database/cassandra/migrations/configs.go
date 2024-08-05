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
	user := viper.GetString("USER_CASSANDRA")
	password := viper.GetString("PASSWORD_CASSANDRA")
	hosts := strings.Split(viper.GetString("HOST_CASSANDRA"), ",")

	if password == "" {
		Conf, err := configs.LoadConfig("./cmd/cassandra/")

		if err != nil {
			return nil, err
		}

		hosts = strings.Split(Conf.HostCassaandra, ",")
		user = Conf.UserCassaandra
		password = Conf.PassCassaandra
	}

	con, err := connection.NewCassandraConnection(hosts, user, password)

	if err != nil {
		return nil, err
	}

	return con, nil

}
