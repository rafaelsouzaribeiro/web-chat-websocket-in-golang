package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
)

func (r *MesssageRepository) getPagination(table string) Pagination {
	s := fmt.Sprintf("SELECT id,page,total FROM %s.%s", entity.KeySpace, table)
	query := r.cql.Query(s)
	iter := query.Iter()
	defer iter.Close()

	var pagination Pagination
	if iter.Scan(&pagination.Id, &pagination.Page, &pagination.Total) {
		return pagination
	}

	return Pagination{Page: 1, Total: 1}

}
