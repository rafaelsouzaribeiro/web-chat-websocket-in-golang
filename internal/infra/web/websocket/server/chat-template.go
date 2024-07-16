package server

import (
	"encoding/base64"
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

	data := struct {
		WebSocketURL string
		Img          string
	}{
		WebSocketURL: fmt.Sprintf("ws://%s:%d%s", server.host, server.port, server.pattern),
		Img:          base64.StdEncoding.EncodeToString(templates.Img),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (server *Server) serveJS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	w.Write([]byte(templates.ChatJS))
}

func (server *Server) serveCSS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Write([]byte(templates.StylesCSS))
}
