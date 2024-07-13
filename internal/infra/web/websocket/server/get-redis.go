package server

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase/dto"
)

func getRedis(startIndex int64, conn *websocket.Conn, key string) {

	totalMessages, err := rdb.LLen(ctx, key).Result()
	if err != nil {
		fmt.Println("Error retrieving total number of messages from Redis:", err)
		return
	}

	if totalMessages > perPage {
		startIndex = totalMessages - perPage
	}

	messages, err := rdb.LRange(ctx, key, startIndex, -1).Result()
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

}
