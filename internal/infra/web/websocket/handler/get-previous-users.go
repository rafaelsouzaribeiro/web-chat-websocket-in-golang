package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase/dto"
)

func (h *MessageHandler) GetUsersFromIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	startIndex, err := strconv.ParseInt(vars["startIndex"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid start index", http.StatusBadRequest)
		return
	}

	StartUIndex = startIndex
	messages, err := h.messageUseCase.ListUsers()

	if err != nil {
		fmt.Printf("error list users %s", err)
	}

	hasMore := StartUIndex > 0

	response := struct {
		Messages []dto.Payload `json:"messages"`
		HasMore  bool          `json:"hasMore"`
	}{
		Messages: *messages,
		HasMore:  hasMore,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
