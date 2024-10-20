package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetInitMessages() (*[]entity.Message, error) {
	pagination := r.getPagination("pagination_messages")

	s := fmt.Sprintf(`select message,pages,username,type,times from %s.messages 
	WHERE pages=?`, entity.KeySpace)
	query := r.cql.Query(s, pagination.Page)
	iter := query.Iter()
	defer iter.Close()

	var message entity.Message
	var messages []entity.Message

	for iter.Scan(&message.Message, &message.Pages,
		&message.Username, &message.Type, &message.Time) {
		messages = append(messages, message)
	}

	return &messages, nil

}
