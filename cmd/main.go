package main

import (
	"log"
	"strconv"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/configs"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/factory"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/web/websocket/handler"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/web/websocket/server"
)

func main() {

	Conf, err := configs.LoadConfig("./")

	if err != nil {
		panic(err)
	}

	port, err := strconv.Atoi(Conf.Port)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	svc := server.NewServer(Conf.HostName, Conf.WsEndPoint, port)
	di := factory.NewFactory(factory.Redis, Conf)
	handler := handler.NewMessageHandler(di)
	svc.Start(handler)

}
