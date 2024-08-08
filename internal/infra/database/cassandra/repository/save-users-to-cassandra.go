package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) SaveUsers(msg *entity.Message) error {
	var save Save
	var total int
	var page int = 1

	s := fmt.Sprintf(`SELECT id,page,total FROM %s.pagination_users`, entity.KeySpace)
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
			page = save.Page
		}
	}

	if strings.TrimSpace(msg.Message) != "" {
		q := fmt.Sprintf(`INSERT INTO %s.users (id, pages, message, username, type, times) 
						  VALUES (?, ?, ?, ?, ?, ?)`, entity.KeySpace)
		err := r.cql.Query(q, uuid.NewString(), page, msg.Message,
			msg.Username, "", time.Now()).Exec()

		if err != nil {
			return err
		}

		if iter.NumRows() == 0 {
			query := fmt.Sprintf(`INSERT INTO %s.pagination_users (id,page,total) VALUES (?,?,?)`,
				entity.KeySpace)

			err = r.cql.Query(query, uuid.NewString(), 1, 1).Exec()

			if err != nil {
				return err
			}
		} else {
			query := fmt.Sprintf(`UPDATE %s.pagination_users SET page = ?, total = ? 
								  WHERE id = ?`, entity.KeySpace)

			err = r.cql.Query(query, page, total, save.Id).Exec()

			if err != nil {
				return err
			}
		}

	}

	return nil
}
