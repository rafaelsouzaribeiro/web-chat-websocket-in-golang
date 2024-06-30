package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase/dto"
)

type Server struct {
	host    string
	port    int
	pattern string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type User struct {
	conn     *websocket.Conn
	username string
	id       string
}

var broadcast = make(chan dto.Payload)
var users = make(map[string]User)
var messageExists = make(map[*websocket.Conn]bool)
var buffer []dto.Payload
var mu sync.Mutex

func NewServer(host, pattern string, port int) *Server {
	return &Server{
		host:    host,
		port:    port,
		pattern: pattern,
	}
}

func (server *Server) ServerWebsocket() {
	http.HandleFunc(server.pattern, handleConnections)
	http.HandleFunc("/chat", serveChat)

	go handleMessages()

	fmt.Printf("Server started on %s:%d \n", server.host, server.port)

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", server.host, server.port), nil)
	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}

func serveChat(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../web/templates/chat.html")
}

func handleMessages() {
	for msg := range broadcast {
		mu.Lock()
		for _, user := range users {
			if err := user.conn.WriteJSON(msg); err != nil {
				fmt.Println("Error sending message:", err)
				user.conn.Close()
				deleteUserByUserName(user.username, false)
			}
		}
		mu.Unlock()
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	mu.Lock()
	for _, v := range buffer {
		conn.WriteJSON(v)
	}
	mu.Unlock()

	defer func() {
		username := getUsernameByConnection(conn)

		mu.Lock()
		delete(messageExists, conn)
		mu.Unlock()

		if username != "" {
			disconnectionMessage := dto.Payload{
				Username: "info",
				Message:  fmt.Sprintf("User %s disconnected", username),
			}

			mu.Lock()
			buffer = append(buffer, disconnectionMessage)
			mu.Unlock()

			broadcast <- disconnectionMessage

			deleteUserByUserName(username, true)
			conn.Close()
		}
	}()

	for {
		var msgs dto.Payload
		err := conn.ReadJSON(&msgs)
		if err != nil {
			break
		}

		if !verifyExistsUser(msgs.Username, conn) {
			if verifyCon(conn, &messageExists) {
				systemMessage := dto.Payload{
					Username: "info",
					Message:  fmt.Sprintf("User already exists: %s", msgs.Username),
				}

				deleteUserByConn(conn, false)
				conn.WriteJSON(systemMessage)
			}
			continue
		}

		if msgs.Type == "message" {
			mu.Lock()
			buffer = append(buffer, msgs)
			mu.Unlock()
			broadcast <- msgs
		} else {
			systemMessage := dto.Payload{
				Username: "info",
				Message:  fmt.Sprintf("User %s connected", msgs.Username),
			}

			mu.Lock()
			id := uuid.New().String()
			users[id] = User{
				conn:     conn,
				username: msgs.Username,
				id:       id,
			}

			buffer = append(buffer, systemMessage)
			mu.Unlock()

			broadcast <- systemMessage
		}
	}
}

func getUsernameByConnection(conn *websocket.Conn) string {
	mu.Lock()
	defer mu.Unlock()
	for _, user := range users {
		if user.conn == conn {
			return user.username
		}
	}
	return ""
}

func deleteUserByUserName(username string, close bool) {
	mu.Lock()
	defer mu.Unlock()
	for k, user := range users {
		if user.username == username {
			if close {
				user.conn.Close()
			}
			delete(users, k)
		}
	}
}

func deleteUserByConn(conn *websocket.Conn, close bool) {
	mu.Lock()
	defer mu.Unlock()
	for k, user := range users {
		if user.conn == conn {
			if close {
				user.conn.Close()
			}
			delete(users, k)
		}
	}
}

func verifyExistsUser(u string, conn *websocket.Conn) bool {
	mu.Lock()
	defer mu.Unlock()
	for _, user := range users {
		if user.conn != conn && u == user.username {
			return false
		}
	}
	return true
}

func verifyCon(s *websocket.Conn, variable *map[*websocket.Conn]bool) bool {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := (*variable)[s]; !exists {
		(*variable)[s] = true
		return true
	}
	return false
}
