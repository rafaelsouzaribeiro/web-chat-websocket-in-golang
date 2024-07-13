package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase/dto"
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
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
				Time:     time.Now(),
			}

			deleteUserByUserName(username, true)

			saveMessageToRedis(disconnectionMessage, "users")
			connected <- disconnectionMessage
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
					Time:     time.Now(),
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

			saveMessageToRedis(msgs, "messages")
			messages <- msgs
		} else {
			systemMessage := dto.Payload{
				Username: "<strong>info</strong>",
				Message:  fmt.Sprintf("User <strong>%s</strong> connected", msgs.Username),
				Time:     time.Now(),
			}

			mu.Lock()
			id := uuid.New().String()
			users[id] = User{
				conn:     conn,
				username: msgs.Username,
				id:       id,
			}

			mu.Unlock()
			getRedis(0, conn, "users")
			getRedis(0, conn, "messages")
			saveMessageToRedis(systemMessage, "users")
			connected <- systemMessage
		}
	}
}
