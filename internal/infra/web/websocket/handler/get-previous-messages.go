package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/usecase/dto"
)

func (h *MessageHandler) GetMessagesFromIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	startIndex, err := strconv.ParseInt(vars["startIndex"], 10, 64)

	if err != nil {
		http.Error(w, "Invalid start index", http.StatusBadRequest)
		return
	}

	StartMIndex = startIndex
	messages, err := h.messageUseCase.ListMessage()

	var ms []dto.Payload
	for _, v := range *messages {
		ms = append(ms, dto.Payload{
			Message:  v.Message,
			Username: v.Username,
			Type:     v.Type,
			Time:     v.Time,
		})
	}

	hasMore := StartMIndex > 0

	response := struct {
		Messages []dto.Payload `json:"messages"`
		HasMore  bool          `json:"hasMore"`
	}{
		Messages: ms,
		HasMore:  hasMore,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
