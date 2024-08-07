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
	hostsDocker := strings.Split(viper.GetString("HOST_CASSANDRA_DOCKER"), ",")

	Conf, err := configs.LoadConfig("./")

	if err != nil {
		panic(err)
	}

	hosts := strings.Split(Conf.HostCassaandra, ",")

	if hostsDocker[0] != "" {
		hosts = hostsDocker
	}

	port, err := strconv.Atoi(Conf.Port)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	svc := server.NewServer(Conf.HostName, Conf.WsEndPoint, port)
	cassandra, err := connection.NewCassandraConnection(hosts, Conf.PassCassaandra, Conf.UserCassaandra)

	if err != nil {
		panic(err)
	}

	di := di.NewMessageCassandraUseCase(cassandra)
	handler := handler.NewMessageHandler(di)
	go svc.Start(handler)
	select {}

}
