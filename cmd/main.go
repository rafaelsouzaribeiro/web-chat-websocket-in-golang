package main

import (
	"log"
	"strconv"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/configs"

	"github.com/spf13/viper"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/web/websocket/server"
)

func main() {

	hostname := viper.GetString("HOST_NAME")
	wsEndpoint := viper.GetString("WS_ENDPOINT")
	portStr := viper.GetString("PORT")

	Conf, err := configs.LoadConfig("../")

	if err != nil {
		panic(err)
	}

	if hostname == "" {
		hostname = Conf.HostName
	}

	if wsEndpoint == "" {
		wsEndpoint = Conf.WsEndPoint
	}

	if portStr == "" {
		portStr = Conf.Port
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	svc := server.NewServer(hostname, wsEndpoint, port)
	go svc.ServerWebsocket()
	select {}

}
