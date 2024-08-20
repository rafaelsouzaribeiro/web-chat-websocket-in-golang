package main

import (
	"log"
	"strconv"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/configs"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/factory"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/di"
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

	f := factory.NewFactory(factory.Cassandra, Conf)
	db, err := f.GetConnection()

	if err != nil {
		panic(err)
	}

	svc := server.NewServer(Conf.HostName, Conf.WsEndPoint, port)
	di := di.NewUseCase(db)
	handler := handler.NewMessageHandler(di)
	svc.Start(handler)

}