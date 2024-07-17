package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/web/templates"
)

func (server *Server) ServerWebsocket() {
	router := mux.NewRouter()
	router.HandleFunc("/chat", server.serveChat).Methods("GET")
	router.HandleFunc("/last-messages/{startIndex}", server.getMessagesFromIndex).Methods("GET")
	router.HandleFunc("/last-users/{startIndex}", server.getUsersFromIndex).Methods("GET")
	router.HandleFunc(server.pattern, handleConnections)
	router.HandleFunc("/js/functions.js", func(w http.ResponseWriter, r *http.Request) {
		server.serveFile(w, "application/javascript", templates.ChatJS)
	})
	router.HandleFunc("/css/styles.css", func(w http.ResponseWriter, r *http.Request) {
		server.serveFile(w, "text/css", templates.StylesCSS)
	})
	router.HandleFunc("/img/background.png", func(w http.ResponseWriter, r *http.Request) {
		server.serveFile(w, "image/png", templates.Img)
	})

	go handleMessages()
	go handleConnected()

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
