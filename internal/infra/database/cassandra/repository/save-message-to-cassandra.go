package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocql/gocql"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) SaveMessage(msg *entity.Message) error {
	var save Save
	var total int
	var page int = 1

	s := fmt.Sprintf(`SELECT id,page,total FROM %s.pagination_messages`, entity.KeySpace)
	query := r.cql.Query(s)
	iter := query.Iter()
	defer iter.Close()

	if iter.Scan(&save.Id, &save.Page, &save.Total) {
		result := save.Total % 20

		if result == 0 {
			total = save.Total + 1
			page = save.Page + 1
		} else {
			total = save.Total + 1
		}
	}

	if strings.TrimSpace(msg.Message) != "" {
		q := fmt.Sprintf(`INSERT INTO %s.messages (id, pages, message, username, type, times) 
						  VALUES (?, ?, ?, ?, ?, ?)`, entity.KeySpace)
		err := r.cql.Query(q, gocql.TimeUUID(), page, msg.Message,
			msg.Username, msg.Type, time.Now()).Exec()

		if err != nil {
			return err
		}

		if iter.NumRows() == 0 {
			query := fmt.Sprintf(`INSERT INTO %s.pagination_messages (id,page,total) VALUES (?,?, ?)`,
				entity.KeySpace)

			err = r.cql.Query(query, gocql.TimeUUID(), 1, 1).Exec()

			if err != nil {
				return err
			}
		} else {
			query := fmt.Sprintf(`UPDATE %s.pagination_messages SET page = ?, total = ? 
								  WHERE id = ?`, entity.KeySpace)

			err = r.cql.Query(query, page, total, save.Id).Exec()

			if err != nil {
				return err
			}
		}

	}

	return nil
}
