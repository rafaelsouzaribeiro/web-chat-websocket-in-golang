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

	hostname := Conf.HostName
	wsEndpoint := Conf.WsEndPoint
	portStr := Conf.Port
	hostRedis := Conf.HostRedis
	portRedis := Conf.PortRedis
	passRedis := Conf.PassRedis

	if hostRedisDocker != "" {
		hostRedis = hostRedisDocker
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	portR, errs := strconv.Atoi(portRedis)
	if errs != nil {
		log.Fatalf("Invalid port: %v", errs)
	}

	svc := server.NewServer(hostname, wsEndpoint, port)
	redis := connection.ConnectingRedis(hostRedis, portR, passRedis)
	di := di.NewMessageRedisUseCase(redis)
	handler := handler.NewMessageHandler(di)
	go svc.Start(handler)
	select {}

}