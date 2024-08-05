package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/configs"

	"github.com/spf13/viper"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/cassandra/connection"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/di"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/web/websocket/handler"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/web/websocket/server"
)

func main() {

	viper.AutomaticEnv()
	hostname := viper.GetString("HOST_NAME")
	wsEndpoint := viper.GetString("WS_ENDPOINT")
	portStr := viper.GetString("PORT")
	user := viper.GetString("USER_CASSANDRA")
	password := viper.GetString("PASSWORD_CASSANDRA")
	hosts := strings.Split(viper.GetString("HOST_CASSANDRA"), ",")

	if hostname == "" {
		Conf, err := configs.LoadConfig("./")

		if err != nil {
			panic(err)
		}

		hostname = Conf.HostName
		wsEndpoint = Conf.WsEndPoint
		portStr = Conf.Port
		user = Conf.UserCassaandra
		password = Conf.PassCassaandra
		hosts = strings.Split(Conf.HostCassaandra, ",")

	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	svc := server.NewServer(hostname, wsEndpoint, port)
	cassandra, err := connection.NewCassandraConnection(hosts, user, password)

	if err != nil {
		panic(err)
	}

	di := di.NewMessageCassandraUseCase(cassandra)
	handler := handler.NewMessageHandler(di)
	go svc.Start(handler)
	select {}

}
