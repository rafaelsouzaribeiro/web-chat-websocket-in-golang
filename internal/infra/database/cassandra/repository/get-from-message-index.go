package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetFromMessageIndex() (*[]entity.Message, error) {

	if entity.PageM == 0 && entity.TotalM == 21 {
		entity.PageM = 2
	}

	if entity.PageM == 1 && (entity.TotalM < 20) {
		return &[]entity.Message{}, nil
	}

	if entity.PointerM == entity.PageM {
		entity.PageM--
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

	entity.PageM--

	return &messages, nil
}
