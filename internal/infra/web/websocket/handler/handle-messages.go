package handler

import "fmt"

func (h *MessageHandler) HandleMessages() {
	for msg := range messages {

		for _, user := range users {
			mu.Lock()
			msg.Username = fmt.Sprintf("<strong>%s</strong>", msg.Username)
			err := user.conn.WriteJSON(msg)
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
}
