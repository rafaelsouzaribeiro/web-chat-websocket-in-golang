package handler

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase/dto"
)

func (h *MessageHandler) run(disconnectionMessage dto.Payload) {
	for _, user := range users {

		mu.Lock()
		err := user.conn.WriteJSON(disconnectionMessage)
		mu.Unlock()

		if err != nil {
			fmt.Println("Error sending message:", err)
			mu.Lock()
			user.conn.Close()
			mu.Unlock()
			h.deleteUserByUserName(user.username, false)
		}

	}
}
