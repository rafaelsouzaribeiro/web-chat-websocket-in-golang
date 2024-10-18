package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *MessageHandler) GetRows(w http.ResponseWriter, r *http.Request) {

	messages, err := h.messageUseCase.GetMessageRows()

	if err != nil {
		fmt.Printf("error reading rows message %s", err)
	}

	users, err := h.messageUseCase.GetUsersRows()

	if err != nil {
		fmt.Printf("error reading rows users %s", err)
	}

	response := struct {
		Messages float64 `json:"rows_messages"`
		Users    float64 `json:"rows_users"`
	}{
		Messages: messages,
		Users:    users,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
