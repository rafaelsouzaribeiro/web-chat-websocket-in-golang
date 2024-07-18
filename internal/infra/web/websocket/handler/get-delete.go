package handler

import "github.com/gorilla/websocket"

func (h *MessageHandler) getUsernameByConnection(conn *websocket.Conn) string {
	mu.Lock()
	defer mu.Unlock()
	for _, user := range users {
		if user.conn == conn {
			return user.username
		}
	}
	return ""
}

func (h *MessageHandler) deleteUserByUserName(username string, close bool) {
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

func (h *MessageHandler) deleteUserByConn(conn *websocket.Conn, close bool) {
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
