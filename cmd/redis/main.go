package main

import (
	"log"
	"strconv"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/configs"

	"github.com/spf13/viper"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/redis/connection"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/di"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/web/websocket/handler"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/web/websocket/server"
)

func main() {

	viper.AutomaticEnv()
	hostRedisDocker := viper.GetString("HOST_REDIS_DOCKER")

	Conf, err := configs.LoadConfig("./")

	if err != nil {
		panic(err)
	}

	hostRedis := Conf.HostRedis

	if hostRedisDocker != "" {
		hostRedis = hostRedisDocker
	}

	port, err := strconv.Atoi(Conf.Port)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	portR, errs := strconv.Atoi(Conf.PortRedis)
	if errs != nil {
		log.Fatalf("Invalid port: %v", errs)
	}

	svc := server.NewServer(Conf.HostName, Conf.WsEndPoint, port)
	redis := connection.ConnectingRedis(hostRedis, portR, Conf.PassRedis)
	di := di.NewMessageRedisUseCase(redis)
	handler := handler.NewMessageHandler(di)
	svc.Start(handler)

}
