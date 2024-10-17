package server

import (
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/configs"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/factory"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/di"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/web/websocket/handler"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/web/websocket/server"
)

func StartServer() {

	Conf, err := configs.LoadConfig("../cmd/")

	if err != nil {
		panic(err)
	}

	port, err := strconv.Atoi(Conf.Port)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	f := factory.NewFactory(factory.Redis, Conf)
	db, err := f.GetConnection()

	if err != nil {
		panic(err)
	}

	di := di.NewUseCase(db)
	svc := server.NewServer(Conf.HostName, Conf.WsEndPoint, port, di)
	handler := handler.NewMessageHandler(di)
	svc.Start(handler)

	time.Sleep(1 * time.Second)

}

var Once sync.Once
