package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase/dto"
)

func (h *MessageHandler) GetMessagesFromIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	startIndex, err := strconv.ParseFloat(vars["startIndex"], 64)

	if err != nil {
		http.Error(w, "Invalid start index", http.StatusBadRequest)
		return
	}

	entity.StartMIndex = startIndex
	messages, err := h.messageUseCase.ListMessage()

	if err != nil {
		fmt.Printf("error list message %s", err)
	}

	response := struct {
		Messages []dto.Payload `json:"messages"`
	}{
		Messages: *messages,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
