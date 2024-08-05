package server

import (
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/configs"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/redis/connection"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/di"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/web/websocket/handler"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/web/websocket/server"
	"github.com/spf13/viper"
)

func StartServer() {
	viper.AutomaticEnv()
	hostname := viper.GetString("HOST_NAME")
	wsEndpoint := viper.GetString("WS_ENDPOINT")
	portStr := viper.GetString("PORT")
	hostRedis := viper.GetString("HOST_REDIS")
	portRedis := viper.GetString("PORT_REDIS")
	passRedis := viper.GetString("PASSWORD_REDIS")

	if hostname == "" {
		Conf, err := configs.LoadConfig("../")

		if err != nil {
			panic(err)
		}

		hostname = Conf.HostName
		wsEndpoint = Conf.WsEndPoint
		portStr = Conf.Port
		hostRedis = Conf.HostRedis
		portRedis = Conf.PortRedis
		passRedis = Conf.PassRedis
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

	time.Sleep(1 * time.Second)

}

var Once sync.Once
