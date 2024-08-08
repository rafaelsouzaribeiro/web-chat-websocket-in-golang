package handler

import "github.com/gorilla/websocket"

func (h *MessageHandler) verifyExistsUser(u string, conn *websocket.Conn) bool {
	mu.Lock()
	defer mu.Unlock()
	for _, user := range users {
		if user.conn != conn && u == user.username {
			return false
		}
	}
	return true
}

func (h *MessageHandler) verifyCon(s *websocket.Conn, variable *map[*websocket.Conn]bool) bool {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := (*variable)[s]; !exists {
		(*variable)[s] = true
		return true
	}
	return false
}

func (h *MessageHandler) verifyId(s *string, id *map[*string]bool) bool {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := (*id)[s]; !exists {
		(*id)[s] = true
		return true
	}
	return false
}
