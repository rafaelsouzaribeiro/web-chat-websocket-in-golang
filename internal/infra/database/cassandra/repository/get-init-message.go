package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetInitMessages() (*[]entity.Message, error) {
	pagination := r.getPagination("pagination_messages")
	var index int
	multi := pagination.Total % int(entity.PerPage)
	index = pagination.Page

	if multi != 0 && pagination.Page != 1 {
		index--
	}

	s := fmt.Sprintf(`select message,pages,username,type,times from %s.messages 
	WHERE pages=?`, entity.KeySpace)
	query := r.cql.Query(s, index)
	iter := query.Iter()
	defer iter.Close()

	var message entity.Message
	var messages []entity.Message

	for iter.Scan(&message.Message, &message.Pages,
		&message.Username, &message.Type, &message.Time) {
		messages = append(messages, message)
	}

	entity.PageM = int64(pagination.Page)
	entity.TotalM = int64(pagination.Total)
	entity.PointerM = int64(index)
	return &messages, nil

}
