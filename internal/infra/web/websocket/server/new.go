package server

import "github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase"

type Server struct {
	host    string
	port    int
	pattern string
	usecase usecase.MessageUsecase
}

func NewServer(host, pattern string, port int, usecase *usecase.MessageUsecase) *Server {
	return &Server{
		host:    host,
		port:    port,
		pattern: pattern,
		usecase: *usecase,
	}
}
