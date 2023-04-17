package scylla

import (
	"log"

	"github.com/luandnh/go-notify/internal/scylla"
)

// CREATE KEYSPACE go_notify WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1};
var RepoClient scylla.IGocqlXClient

func InitRepo() {
	if err := RepoClient.GetSession().ExecStmt(`CREATE TYPE IF NOT EXISTS extra_value (
		key text,
		value text,
		label text
	)`); err != nil {
		log.Fatalf("Error creating type: %v", err)
	}
}
