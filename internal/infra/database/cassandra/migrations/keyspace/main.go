package main

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/entity"
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/internal/infra/database/cassandra/migrations"
)

var cql = ""

func main() {

	con, err := migrations.SetVariables()

	if err != nil {
		panic(err)
	}

	defer con.Close()

	setCommands()

	err = con.Query(cql).Exec()

	if err != nil {
		panic(err)
	}

}

func setCommands() {
	cql = fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS  %s
				WITH REPLICATION = {
				'class' : 'SimpleStrategy',
				'replication_factor' : 1 
				};`, entity.KeySpace)
}
