package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase/dto"
)

func (server *Server) getUsersFromIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	startIndex, err := strconv.ParseInt(vars["startIndex"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid start index", http.StatusBadRequest)
		return
	}

	endIndex := startIndex - 1
	startIndex = endIndex - 19
	if startIndex < 0 {
		startIndex = 0
	}

	messages, err := rdb.LRange(ctx, "users", startIndex, endIndex).Result()
	if err != nil {
		http.Error(w, "Error retrieving messages from Redis", http.StatusInternalServerError)
		return
	}

	var payloads []dto.Payload
	for _, msg := range messages {
		var payload dto.Payload
		if err := json.Unmarshal([]byte(msg), &payload); err == nil {
			payloads = append(payloads, payload)
		}
	}

	hasMore := startIndex > 0

	response := struct {
		Messages []dto.Payload `json:"messages"`
		HasMore  bool          `json:"hasMore"`
	}{
		Messages: payloads,
		HasMore:  hasMore,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
