package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetFromUsersIndex() (*[]entity.Message, error) {

	entity.PageU--

	if entity.PageU == 0 && entity.TotalU <= 21 {
		entity.PageU = 2
	}

	if entity.PageU == 1 && entity.TotalU == 21 {
		return &[]entity.Message{}, nil
	}

	s := fmt.Sprintf(`select message,pages,username,type,times from %s.users 
	WHERE pages=? ORDER BY times DESC`, entity.KeySpace)
	query := r.cql.Query(s, entity.PageU)
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
