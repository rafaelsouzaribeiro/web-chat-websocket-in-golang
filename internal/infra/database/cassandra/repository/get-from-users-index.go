package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetFromUsersIndex() (*[]entity.Message, error) {

	s := fmt.Sprintf(`select message,pages,username,type,times from %s.users 
	WHERE pages=? ORDER BY times DESC`, entity.KeySpace)
	query := r.cql.Query(s, entity.StartUIndex)
	iter := query.Iter()
	defer iter.Close()

	var message entity.Message
	var messages []entity.Message

	for iter.Scan(&message.Message, &message.Pages,
		&message.Username, &message.Type, &message.Time) {
		messages = append(messages, message)
	}

	entity.StartUIndex++

	return &messages, nil
}
