package main

import (
	"log"
	"os"
	"strconv"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/web/websocket/server"
)

func main() {

	appName := os.Getenv("APP_NAME")
	wsEndpoint := os.Getenv("WS_ENDPOINT")
	portStr := os.Getenv("PORT")

	if appName == "" {
		appName = "localhost"
	}

	if wsEndpoint == "" {
		wsEndpoint = "/ws"
	}

	if portStr == "" {
		portStr = "8080"
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	svc := server.NewServer(appName, wsEndpoint, port)
	go svc.ServerWebsocket()
	select {}

}
