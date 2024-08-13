package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetInitUsers() (*[]entity.Message, error) {
	pagination := r.getPagination("pagination_messages")

	multi := pagination.Total % 20

	if multi != 0 && pagination.Page != 1 {
		pagination.Page--
	}

	s := fmt.Sprintf(`select message,pages,username,type,times from %s.users 
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

	entity.PageU = int64(pagination.Page)

	return &messages, nil

}
