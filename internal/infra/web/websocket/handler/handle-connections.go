package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase/dto"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *MessageHandler) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		username := h.getUsernameByConnection(conn)

		mu.Lock()
		delete(messageExists, conn)
		mu.Unlock()

		if username != "" {
			disconnectionMessage := dto.Payload{
				Username: "<strong>info</strong>",
				Message:  fmt.Sprintf("User <strong>%s</strong> disconnected", username),
				Time:     time.Now(),
			}

			h.deleteUserByUserName(username, true)
			h.run(disconnectionMessage)
			h.messageUseCase.SaveUsers(disconnectionMessage)
			conn.Close()
		}
	}()

	for {
		var msgs dto.Payload
		err := conn.ReadJSON(&msgs)
		if err != nil {
			break
		}

		if !h.verifyExistsUser(msgs.Username, conn) {
			if h.verifyCon(conn, &messageExists) {
				systemMessage := dto.Payload{
					Username: "<strong>info</strong>",
					Message:  fmt.Sprintf("User already exists: <strong>%s</strong>", msgs.Username),
					Time:     time.Now(),
				}

				h.deleteUserByConn(conn, false)
				conn.WriteJSON(systemMessage)

			}
			continue
		}

		if msgs.Type == "message" {
			mu.Lock()
			msgs.Username = fmt.Sprintf("<strong>%s</strong>", msgs.Username)
			h.messageUseCase.SaveMessage(msgs)
			mu.Unlock()
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
			message, err := h.messageUseCase.GetInitusers()

			if err != nil {
				fmt.Printf("Error %v \n", err)
			}

			if message != nil {
				for _, payload := range *message {
					mu.Lock()
					conn.WriteJSON(payload)
					mu.Unlock()
				}
			}

			messa, err := h.messageUseCase.GetInitMessages()

			if err != nil {
				fmt.Printf("Error %v \n", err)
			}

			if messa != nil {
				for _, payload := range *messa {
					mu.Lock()
					conn.WriteJSON(payload)
					mu.Unlock()
				}
			}

			mu.Lock()
			h.messageUseCase.SaveUsers(systemMessage)
			mu.Unlock()
			connected <- systemMessage
		}
	}
}
