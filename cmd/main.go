package main

import (
	"github.com/spf13/viper"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/web/websocket/server"
)

func main() {

	appName := viper.GetString("APP_NAME")
	wsEndpoint := viper.GetString("WS_ENDPOINT")
	portStr := viper.GetInt("PORT")

	if appName == "" {
		appName = "localhost"
	}

	if wsEndpoint == "" {
		wsEndpoint = "/ws"
	}

	if portStr == 0 {
		portStr = 8080
	}

	svc := server.NewServer(appName, wsEndpoint, portStr)
	go svc.ServerWebsocket()
	select {}

}
