package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase/dto"
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	messages, err := rdb.LRange(ctx, "chat_messages", 0, -1).Result()
	if err != nil {
		fmt.Println("Error retrieving messages from Redis:", err)
	} else {
		for _, msg := range messages {
			var payload dto.Payload
			if err := json.Unmarshal([]byte(msg), &payload); err == nil {
				conn.WriteJSON(payload)
			}
		}
	}

	defer func() {
		username := getUsernameByConnection(conn)

		mu.Lock()
		delete(messageExists, conn)
		mu.Unlock()

		if username != "" {
			disconnectionMessage := dto.Payload{
				Username: "<strong>info</strong>",
				Message:  fmt.Sprintf("User <strong>%s</strong> disconnected", username),
			}

			deleteUserByUserName(username, true)

			saveMessageToRedis(disconnectionMessage)
			broadcast <- disconnectionMessage
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
					Username: "<strong>info</strong>",
					Message:  fmt.Sprintf("User already exists: <strong>%s</strong>", msgs.Username),
				}

				deleteUserByConn(conn, false)
				conn.WriteJSON(systemMessage)
			}
			continue
		}

		if msgs.Type == "message" {
			mu.Lock()
			msgs.Username = fmt.Sprintf("<strong>%s</strong>", msgs.Username)
			mu.Unlock()

			saveMessageToRedis(msgs)
			broadcast <- msgs
		} else {
			systemMessage := dto.Payload{
				Username: "<strong>info</strong>",
				Message:  fmt.Sprintf("User <strong>%s</strong> connected", msgs.Username),
			}

			mu.Lock()
			id := uuid.New().String()
			users[id] = User{
				conn:     conn,
				username: msgs.Username,
				id:       id,
			}

			mu.Unlock()

			saveMessageToRedis(systemMessage)
			broadcast <- systemMessage
		}
	}
}
