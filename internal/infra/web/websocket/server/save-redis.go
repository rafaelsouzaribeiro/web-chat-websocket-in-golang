package server

import (
	"encoding/json"
	"fmt"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase/dto"
)

func saveMessageToRedis(msg dto.Payload) {
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error marshaling message:", err)
		return
	}

	err = rdb.RPush(ctx, "chat_messages", data).Err()
	if err != nil {
		fmt.Println("Error saving message to Redis:", err)
	}
}
