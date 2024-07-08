package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (server *Server) ServerWebsocket() {
	router := mux.NewRouter()
	router.HandleFunc("/chat", server.serveChat).Methods("GET")
	router.HandleFunc("/last-messages/{startIndex}", server.getMessagesFromIndex).Methods("GET")

	router.HandleFunc(server.pattern, handleConnections)

	go handleMessages()

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", server.host, server.port),
		Handler: router,
	}

	go func() {
		fmt.Printf("Server started on %s:%d \n", server.host, server.port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %v\n", err)
		}
	}()

	gracefulShutdown(httpServer)
}
