package main

import (
	"log"
	"strconv"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/configs"

	"github.com/spf13/viper"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/web/websocket/server"
)

func main() {

	viper.AutomaticEnv()
	hostname := viper.GetString("HOST_NAME")
	wsEndpoint := viper.GetString("WS_ENDPOINT")
	portStr := viper.GetString("PORT")
	hostRedis := viper.GetString("HOST_REDIS")
	portRedis := viper.GetString("PORT_REDIS")

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
	svc.ConnectingRedis(hostRedis, portR)
	go svc.ServerWebsocket()
	select {}

}
