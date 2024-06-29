package main

import "github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/web/websocket/server"

func main() {

	svc := server.NewServer("localhost", "/ws", 8080)
	go svc.ServerWebsocket()
	select {}

}
