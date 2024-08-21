package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

var condM bool = false

func (r *MesssageRepository) GetFromMessageIndex() (*[]entity.Message, error) {

	if entity.PageM == 1 && !condM {
		entity.PageM = 2
		condM = true
	}

	entity.PageM--

	if entity.PageM == 1 && condM {
		return &[]entity.Message{}, nil
	}

	s := fmt.Sprintf(`select message,pages,username,type,times from %s.messages 
					  WHERE pages=?`, entity.KeySpace)
	query := r.cql.Query(s, entity.PageM)
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
