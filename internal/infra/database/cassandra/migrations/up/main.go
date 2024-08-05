package main

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/cassandra/migrations"
)

var cql = make([]string, 4)

func main() {

	con, err := migrations.SetVariables()

	if err != nil {
		panic(err)
	}

	defer con.Close()

	setCommands()

	for _, v := range cql {
		err = con.Query(v).Exec()

		if err != nil {
			panic(err)
		}
	}

}

func setCommands() {
	cql[0] = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.messages (
            id TEXT,
            message TEXT,
            pages INT,
            username TEXT,
            type TEXT,
            times TIMESTAMP,
            PRIMARY KEY (pages,times )
        ) WITH CLUSTERING ORDER BY (times DESC);`, entity.KeySpace)

	cql[1] = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.pagination_users (
		id TEXT PRIMARY KEY,
		page INT,
		total INT);`, entity.KeySpace)

	cql[2] = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.pagination_messages (
    id TEXT PRIMARY KEY,
    page INT,
    total INT
	);`, entity.KeySpace)

	cql[3] = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.users (
				id TEXT,
				message TEXT,
				pages INT,
				username TEXT,
				type TEXT,
				times TIMESTAMP,
				PRIMARY KEY (pages,times )
			) WITH CLUSTERING ORDER BY (times ASC);`, entity.KeySpace)
}
