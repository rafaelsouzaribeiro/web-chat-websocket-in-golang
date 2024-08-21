package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

var condU bool = false

func (r *MesssageRepository) GetFromUsersIndex() (*[]entity.Message, error) {

	if entity.PageU == 1 && !condU {
		entity.PageU = 2
		condM = true
	}

	entity.PageU--

	if entity.PageU == 1 && condU {
		return &[]entity.Message{}, nil
	}

	//println(">>", entity.StartUIndex, entity.PageU)
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
