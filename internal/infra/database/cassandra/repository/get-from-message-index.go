package repository

import (
	"fmt"
	"time"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) GetFromMessageIndex() (*[]entity.Message, error) {

	Once.Do(func() { entity.StartMIndex-- })

	if entity.StartMIndex == entity.PageM {
		entity.StartMIndex++
	}

	tenMinutesAgo := time.Now().Add(-20 * time.Minute)
	s := fmt.Sprintf(`select message,pages,username,type,times from %s.messages 
					  WHERE pages=? AND times < ? ALLOW FILTERING`, entity.KeySpace)
	query := r.cql.Query(s, entity.StartMIndex, tenMinutesAgo)
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
