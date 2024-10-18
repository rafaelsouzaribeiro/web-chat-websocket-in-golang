package server

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/web/templates"
)

func (server *Server) serveChat(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").Parse(templates.Chat)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	messages, err := server.usecase.GetMessageRows()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users, err := server.usecase.GetUsersRows()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		WebSocketURL string
		indexm       float64
		indexU       float64
	}{
		WebSocketURL: fmt.Sprintf("ws://%s:%d%s", server.host, server.port, server.pattern),
		indexm:       messages,
		indexU:       users,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (server *Server) serveFile(w http.ResponseWriter, contentType, filePath string) {
	w.Header().Set("Content-Type", contentType)
	w.Write([]byte(filePath))
}
