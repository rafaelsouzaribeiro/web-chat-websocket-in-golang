package main

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/cassandra/migrations"
)

var cql = make([]string, 5)

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
            id TIMEUUID,
            message TEXT,
            pages INT,
            username TEXT,
            type TEXT,
            times TIMESTAMP,
            PRIMARY KEY (pages,times )
        ) WITH CLUSTERING ORDER BY (times ASC);`, entity.KeySpace)

	cql[1] = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.pagination_users (
		page INT,
		total INT,
		PRIMARY KEY (page,total));`, entity.KeySpace)

	cql[2] = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.pagination_users (
		page INT,
		total INT,
		PRIMARY KEY (page,total));`, entity.KeySpace)

	cql[3] = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.pagination_messages (
		page INT,
		total INT,
		PRIMARY KEY (page,total));`, entity.KeySpace)

	cql[4] = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.users (
				id TIMEUUID,
				message TEXT,
				pages INT,
				username TEXT,
				type TEXT,
				times TIMESTAMP,
				PRIMARY KEY (pages,times )
			) WITH CLUSTERING ORDER BY (times ASC);`, entity.KeySpace)
}
