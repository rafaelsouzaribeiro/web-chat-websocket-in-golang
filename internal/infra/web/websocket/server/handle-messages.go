package server

import "fmt"

func handleMessages() {
	for msg := range broadcast {
		mu.Lock()
		for _, user := range users {
			msg.Username = fmt.Sprintf("<strong>%s</strong>", msg.Username)

			if err := user.conn.WriteJSON(msg); err != nil {
				fmt.Println("Error sending message:", err)
				user.conn.Close()
				deleteUserByUserName(user.username, false)
			}
		}
		mu.Unlock()
	}
}
