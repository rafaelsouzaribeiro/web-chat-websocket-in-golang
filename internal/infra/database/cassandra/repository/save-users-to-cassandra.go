package repository

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/gocql/gocql"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) SaveUsers(msg *entity.Message) error {
	var save Save
	var total int
	var page int = 1

	s := fmt.Sprintf(`SELECT page,total FROM %s.pagination_users`, entity.KeySpace)
	query := r.cql.Query(s)
	iter := query.Iter()
	defer iter.Close()

	if iter.Scan(&save.Page, &save.Total) {
		result := float64(save.Page / 20)

		if math.Mod(result, 1) == 0 {
			total = save.Total + 1

		} else {
			total = save.Total + 1
			page = save.Page + 1
		}
	}

	if strings.TrimSpace(msg.Message) != "" {
		q := fmt.Sprintf(`INSERT INTO %s.users (id, pages, message, username, type, times) 
						  VALUES (?, ?, ?, ?, ?, ?)`, entity.KeySpace)
		err := r.cql.Query(q, gocql.TimeUUID(), page, msg.Message,
			msg.Username, "", time.Now()).Exec()

		if err != nil {
			return err
		}

		if iter.NumRows() == 0 {
			query := fmt.Sprintf(`INSERT INTO %s.pagination_users (page,total) VALUES (?, ?)`,
				entity.KeySpace)

			err = r.cql.Query(query, 1, 1).Exec()

			if err != nil {
				return err
			}
		} else {
			query := fmt.Sprintf(`UPDATE %s.pagination_users SET page = ?, total = ? 
								  WHERE page = ?`, entity.KeySpace)

			err = r.cql.Query(query, page, total, save.Page).Exec()

			if err != nil {
				return err
			}
		}

	}

	return nil
}
